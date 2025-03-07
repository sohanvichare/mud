import { awaitStreamValue, DoWork, filterNullish, keccak256, streamToDefinedComputed } from "@latticexyz/utils";
import {
  bufferTime,
  catchError,
  concat,
  concatMap,
  filter,
  ignoreElements,
  map,
  Observable,
  of,
  Subject,
  take,
} from "rxjs";
import {
  isNetworkComponentUpdateEvent,
  NetworkComponentUpdate,
  NetworkEvent,
  NetworkEvents,
  SyncStateStruct,
  SyncWorkerConfig,
} from "../types";
import { Components, ComponentValue, SchemaOf } from "@latticexyz/recs";
import {
  createCacheStore,
  getCacheStoreEntries,
  getIndexDBCacheStoreBlockNumber,
  getIndexDbECSCache,
  loadIndexDbCacheStore,
  saveCacheStoreToIndexDb,
  storeEvent,
  storeEvents,
} from "./CacheStore";
import { createReconnectingProvider } from "../createProvider";
import { computed } from "mobx";
import {
  createSnapshotClient,
  createDecode,
  createFetchWorldEventsInBlockRange,
  getSnapshotBlockNumber,
  fetchSnapshotChunked,
  createLatestEventStreamRPC,
  createLatestEventStreamService,
  createTransformWorldEventsFromStream,
  createFetchSystemCallsFromEvents,
  fetchEventsInBlockRangeChunked,
} from "./syncUtils";
import { createBlockNumberStream } from "../createBlockNumberStream";
import { GodID, SyncState } from "./constants";

export enum InputType {
  Ack,
  Config,
}
export type Config = { type: InputType.Config; data: SyncWorkerConfig };
export type Ack = { type: InputType.Ack };
export const ack = { type: InputType.Ack as const };
export type Input = Config | Ack;

export class SyncWorker<C extends Components> implements DoWork<Input, NetworkEvent<C>[]> {
  private input$ = new Subject<Input>();
  private output$ = new Subject<NetworkEvent<C>>();
  private syncState: SyncStateStruct = { state: SyncState.CONNECTING, msg: "", percentage: 0 };

  constructor() {
    this.init();
  }

  /**
   * Pass a loading state component update to the main thread.
   * Can be used to indicate the initial loading state on a loading screen.
   * @param loadingState {
   *  state: {@link SyncState},
   *  msg: Message to describe the current loading step.
   *  percentage: Number between 0 and 100 to describe the loading progress.
   * }
   * @param blockNumber Optional: block number to pass in the component update.
   */
  private setLoadingState(
    loadingState: Partial<{ state: SyncState; msg: string; percentage: number }>,
    blockNumber = 0
  ) {
    const newLoadingState = { ...this.syncState, ...loadingState };
    this.syncState = newLoadingState;
    const update: NetworkComponentUpdate<C> = {
      type: NetworkEvents.NetworkComponentUpdate,
      component: keccak256("component.LoadingState"),
      value: newLoadingState as unknown as ComponentValue<SchemaOf<C[keyof C]>>,
      entity: GodID,
      txHash: "worker",
      lastEventInTx: false,
      blockNumber,
    };

    this.output$.next(update);
  }

  /**
   * Start the sync process.
   *
   * 1. Get config
   * 2. Load initial state
   *   2.1 Get cache block number
   *   2.2 Get snapshot block number
   *   2.3 Load from more recent source
   * 3. Cach up to current block number by requesting events from RPC ( -> TODO: Replace with own service)
   * 4. Keep in sync
   *  4.1 If available keep in sync with streaming service
   *  4.2 Else keep in sync with RPC
   */
  private async init() {
    this.setLoadingState({ state: SyncState.CONNECTING, msg: "Connecting...", percentage: 0 });

    // Turn config into variable accessible outside the stream
    const computedConfig = await streamToDefinedComputed(
      this.input$.pipe(
        map((e) => (e.type === InputType.Config ? e.data : undefined)),
        filterNullish()
      )
    );
    const config = computedConfig.get();
    const {
      snapshotServiceUrl,
      streamServiceUrl,
      chainId,
      worldContract,
      provider: { options: providerOptions },
      initialBlockNumber,
      fetchSystemCalls,
      disableCache,
    } = config;

    // Set default values for cacheAgeThreshold and cacheInterval
    const cacheAgeThreshold = config.cacheAgeThreshold || 100;
    const cacheInterval = config.cacheInterval || 1;

    // Set up
    const { providers } = await createReconnectingProvider(computed(() => computedConfig.get().provider));
    const provider = providers.get().json;
    const snapshotClient = snapshotServiceUrl ? createSnapshotClient(snapshotServiceUrl) : undefined;
    const indexDbCache = await getIndexDbECSCache(chainId, worldContract.address);
    const decode = createDecode(worldContract, provider);
    const fetchWorldEvents = createFetchWorldEventsInBlockRange(
      provider,
      worldContract,
      providerOptions?.batch,
      decode
    );
    const transformWorldEvents = createTransformWorldEventsFromStream(decode);

    // Start syncing current events, but only start streaming to output once gap between initial state and current block is closed

    console.log("[SyncWorker] start initial sync");
    this.setLoadingState({ state: SyncState.INITIAL, msg: "Starting initial sync", percentage: 0 });
    let passLiveEventsToOutput = false;
    const cacheStore = { current: createCacheStore() };
    const { blockNumber$ } = createBlockNumberStream(providers);
    // The RPC is only queried if this stream is subscribed to
    const latestEventRPC$ = createLatestEventStreamRPC(
      blockNumber$,
      fetchWorldEvents,
      fetchSystemCalls ? createFetchSystemCallsFromEvents(provider) : undefined
    );
    const latestEvent$ = streamServiceUrl
      ? createLatestEventStreamService(streamServiceUrl, worldContract.address, transformWorldEvents).pipe(
          catchError((err) => {
            console.warn("[SyncWorker || via Stream Service]", err);
            console.info("[SyncWorker || falling back to RPC]", err);
            return latestEventRPC$;
          })
        )
      : latestEventRPC$;

    const initialLiveEvents: NetworkComponentUpdate<Components>[] = [];
    latestEvent$.subscribe((event) => {
      // If initial sync is in progress, temporary store the events to apply later
      // Ignore system calls during initial sync
      if (!passLiveEventsToOutput) {
        if (isNetworkComponentUpdateEvent(event)) initialLiveEvents.push(event);
        return;
      }

      if (isNetworkComponentUpdateEvent(event)) {
        // Store cache to indexdb every block
        if (event.blockNumber > cacheStore.current.blockNumber + 1 && event.blockNumber % cacheInterval === 0)
          saveCacheStoreToIndexDb(indexDbCache, cacheStore.current);

        storeEvent(cacheStore.current, event);
      }

      this.output$.next(event as NetworkEvent<C>);
    });
    const streamStartBlockNumberPromise = awaitStreamValue(blockNumber$);

    // Load initial state from cache or snapshot service
    this.setLoadingState({ state: SyncState.INITIAL, msg: "Fetching cache block number", percentage: 0 });
    const cacheBlockNumber = !disableCache ? await getIndexDBCacheStoreBlockNumber(indexDbCache) : -1;
    this.setLoadingState({ percentage: 50 });
    const snapshotBlockNumber = await getSnapshotBlockNumber(snapshotClient, worldContract.address);
    this.setLoadingState({ percentage: 100 });
    console.log(
      `[SyncWorker] cache block: ${cacheBlockNumber}, snapshot block: ${
        snapshotBlockNumber > 0 ? snapshotBlockNumber : "Unavailable"
      }, start sync at ${initialBlockNumber}`
    );

    let initialState = createCacheStore();
    if (initialBlockNumber > Math.max(cacheBlockNumber, snapshotBlockNumber)) {
      initialState.blockNumber = initialBlockNumber;
    } else {
      // Load from cache if the snapshot is less than <cacheAgeThreshold> blocks newer than the cache
      const syncFromSnapshot = snapshotClient && snapshotBlockNumber > cacheBlockNumber + cacheAgeThreshold;

      if (syncFromSnapshot) {
        this.setLoadingState({ state: SyncState.INITIAL, msg: "Fetching initial state from snapshot", percentage: 0 });
        initialState = await fetchSnapshotChunked(
          snapshotClient,
          worldContract.address,
          decode,
          config.snapshotNumChunks,
          (percentage: number) => this.setLoadingState({ percentage }),
          config.pruneOptions
        );
      } else {
        this.setLoadingState({ state: SyncState.INITIAL, msg: "Fetching initial state from cache", percentage: 0 });
        initialState = await loadIndexDbCacheStore(indexDbCache);
        this.setLoadingState({ percentage: 100 });
      }

      console.log(`[SyncWorker] got ${initialState.state.size} items from ${syncFromSnapshot ? "snapshot" : "cache"}`);
    }

    // Load events from gap between initial state and current block number from RPC
    const streamStartBlockNumber = await streamStartBlockNumberPromise;
    this.setLoadingState({
      state: SyncState.INITIAL,
      msg: `Fetching state from block ${initialState.blockNumber} to ${streamStartBlockNumber}`,
      percentage: 0,
    });

    const gapStateEvents = await fetchEventsInBlockRangeChunked(
      fetchWorldEvents,
      initialState.blockNumber,
      streamStartBlockNumber,
      50,
      (percentage: number) => this.setLoadingState({ percentage })
    );

    console.log(
      `[SyncWorker || via JSON-RPC] got ${gapStateEvents.length} items from block range ${initialState.blockNumber} -> ${streamStartBlockNumber}`
    );

    // Merge initial state, gap state and live events since initial sync started
    storeEvents(initialState, [...gapStateEvents, ...initialLiveEvents]);
    cacheStore.current = initialState;
    console.log(`[SyncWorker] initial sync state size: ${cacheStore.current.state.size}`);

    this.setLoadingState({
      state: SyncState.INITIAL,
      msg: `Initializing with ${cacheStore.current.state.size} state entries`,
      percentage: 0,
    });

    // Pass current cacheStore to output and start passing live events
    let i = 0;
    for (const update of getCacheStoreEntries(cacheStore.current)) {
      i++;
      this.output$.next(update as NetworkEvent<C>);
      if (i % 5000 === 0) {
        const percentage = Math.floor((i / cacheStore.current.state.size) * 100);
        this.setLoadingState({ percentage });
      }
    }

    // Save initial state to cache
    saveCacheStoreToIndexDb(indexDbCache, cacheStore.current);

    // Let the client know loading is complete
    this.setLoadingState(
      { state: SyncState.LIVE, msg: `Streaming live events`, percentage: 100 },
      cacheStore.current.blockNumber
    );
    passLiveEventsToOutput = true;
  }

  public work(input$: Observable<Input>): Observable<NetworkEvent<C>[]> {
    input$.subscribe(this.input$);
    const throttledOutput$ = new Subject<NetworkEvent<C>[]>();

    this.output$
      .pipe(
        bufferTime(16, null, 50),
        filter((updates) => updates.length > 0),
        concatMap((updates) =>
          concat(
            of(updates),
            input$.pipe(
              filter((e) => e.type === InputType.Ack),
              take(1),
              ignoreElements()
            )
          )
        )
      )
      .subscribe(throttledOutput$);

    return throttledOutput$;
  }
}
