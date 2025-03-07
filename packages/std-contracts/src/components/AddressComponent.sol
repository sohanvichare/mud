// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;
import "solecs/Component.sol";

contract AddressComponent is Component {
  constructor(address world, uint256 id) Component(world, id) {}

  function getSchema() public pure override returns (string[] memory keys, LibTypes.SchemaValue[] memory values) {
    keys = new string[](1);
    values = new LibTypes.SchemaValue[](1);

    keys[0] = "value";
    values[0] = LibTypes.SchemaValue.UINT256;
  }

  function set(uint256 entity, address value) public {
    set(entity, abi.encode(value));
  }

  function getValue(uint256 entity) public view returns (address) {
    address value = abi.decode(getRawValue(entity), (address));
    return value;
  }

  function getEntitiesWithValue(address value) public view returns (uint256[] memory) {
    return getEntitiesWithValue(abi.encode(value));
  }
}
