// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: proto/ecs-relay.proto

package ecs_relay

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Identifies a client connecting to Relay Service.
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Signature that a client must provide to prove ownership of identity.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature    *Signature    `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Subscription *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionRequest) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SubscriptionRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{4}
}

func (x *Subscription) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label   string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{5}
}

func (x *PushRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PushRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature *Signature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Label     string     `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Messages  []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PushManyRequest) Reset() {
	*x = PushManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushManyRequest) ProtoMessage() {}

func (x *PushManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushManyRequest.ProtoReflect.Descriptor instead.
func (*PushManyRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{6}
}

func (x *PushManyRequest) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *PushManyRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PushManyRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{7}
}

type CountIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountIdentitiesRequest) Reset() {
	*x = CountIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountIdentitiesRequest) ProtoMessage() {}

func (x *CountIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountIdentitiesRequest.ProtoReflect.Descriptor instead.
func (*CountIdentitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{8}
}

type CountIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountIdentitiesResponse) Reset() {
	*x = CountIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountIdentitiesResponse) ProtoMessage() {}

func (x *CountIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountIdentitiesResponse.ProtoReflect.Descriptor instead.
func (*CountIdentitiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{9}
}

func (x *CountIdentitiesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{10}
}

type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{11}
}

func (x *BalanceResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_proto_ecs_relay_proto protoreflect.FileDescriptor

var file_proto_ecs_relay_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x22, 0x1e, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x29, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x83, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0x50, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2d,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x0e, 0x0a,
	0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a,
	0x16, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xfa, 0x05, 0x0a, 0x0f, 0x45, 0x43, 0x53, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x65, 0x63,
	0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x1a, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x12, 0x13, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5b,
	0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x20, 0x2e,
	0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x73, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x13, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x11, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0a,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x65, 0x63, 0x73,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x37, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x75, 0x73, 0x68, 0x12, 0x18, 0x2e,
	0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ecs_relay_proto_rawDescOnce sync.Once
	file_proto_ecs_relay_proto_rawDescData = file_proto_ecs_relay_proto_rawDesc
)

func file_proto_ecs_relay_proto_rawDescGZIP() []byte {
	file_proto_ecs_relay_proto_rawDescOnce.Do(func() {
		file_proto_ecs_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ecs_relay_proto_rawDescData)
	})
	return file_proto_ecs_relay_proto_rawDescData
}

var file_proto_ecs_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_ecs_relay_proto_goTypes = []interface{}{
	(*Identity)(nil),                // 0: ecsrelay.Identity
	(*Signature)(nil),               // 1: ecsrelay.Signature
	(*Message)(nil),                 // 2: ecsrelay.Message
	(*SubscriptionRequest)(nil),     // 3: ecsrelay.SubscriptionRequest
	(*Subscription)(nil),            // 4: ecsrelay.Subscription
	(*PushRequest)(nil),             // 5: ecsrelay.PushRequest
	(*PushManyRequest)(nil),         // 6: ecsrelay.PushManyRequest
	(*PushResponse)(nil),            // 7: ecsrelay.PushResponse
	(*CountIdentitiesRequest)(nil),  // 8: ecsrelay.CountIdentitiesRequest
	(*CountIdentitiesResponse)(nil), // 9: ecsrelay.CountIdentitiesResponse
	(*BalanceRequest)(nil),          // 10: ecsrelay.BalanceRequest
	(*BalanceResponse)(nil),         // 11: ecsrelay.BalanceResponse
}
var file_proto_ecs_relay_proto_depIdxs = []int32{
	1,  // 0: ecsrelay.SubscriptionRequest.signature:type_name -> ecsrelay.Signature
	4,  // 1: ecsrelay.SubscriptionRequest.subscription:type_name -> ecsrelay.Subscription
	2,  // 2: ecsrelay.PushRequest.message:type_name -> ecsrelay.Message
	1,  // 3: ecsrelay.PushManyRequest.signature:type_name -> ecsrelay.Signature
	2,  // 4: ecsrelay.PushManyRequest.messages:type_name -> ecsrelay.Message
	1,  // 5: ecsrelay.ECSRelayService.Authenticate:input_type -> ecsrelay.Signature
	1,  // 6: ecsrelay.ECSRelayService.Revoke:input_type -> ecsrelay.Signature
	1,  // 7: ecsrelay.ECSRelayService.Ping:input_type -> ecsrelay.Signature
	8,  // 8: ecsrelay.ECSRelayService.CountAuthenticated:input_type -> ecsrelay.CountIdentitiesRequest
	8,  // 9: ecsrelay.ECSRelayService.CountConnected:input_type -> ecsrelay.CountIdentitiesRequest
	3,  // 10: ecsrelay.ECSRelayService.Subscribe:input_type -> ecsrelay.SubscriptionRequest
	3,  // 11: ecsrelay.ECSRelayService.Unsubscribe:input_type -> ecsrelay.SubscriptionRequest
	1,  // 12: ecsrelay.ECSRelayService.OpenStream:input_type -> ecsrelay.Signature
	5,  // 13: ecsrelay.ECSRelayService.PushStream:input_type -> ecsrelay.PushRequest
	5,  // 14: ecsrelay.ECSRelayService.Push:input_type -> ecsrelay.PushRequest
	10, // 15: ecsrelay.ECSRelayService.MinBalanceForPush:input_type -> ecsrelay.BalanceRequest
	0,  // 16: ecsrelay.ECSRelayService.Authenticate:output_type -> ecsrelay.Identity
	0,  // 17: ecsrelay.ECSRelayService.Revoke:output_type -> ecsrelay.Identity
	0,  // 18: ecsrelay.ECSRelayService.Ping:output_type -> ecsrelay.Identity
	9,  // 19: ecsrelay.ECSRelayService.CountAuthenticated:output_type -> ecsrelay.CountIdentitiesResponse
	9,  // 20: ecsrelay.ECSRelayService.CountConnected:output_type -> ecsrelay.CountIdentitiesResponse
	4,  // 21: ecsrelay.ECSRelayService.Subscribe:output_type -> ecsrelay.Subscription
	4,  // 22: ecsrelay.ECSRelayService.Unsubscribe:output_type -> ecsrelay.Subscription
	2,  // 23: ecsrelay.ECSRelayService.OpenStream:output_type -> ecsrelay.Message
	7,  // 24: ecsrelay.ECSRelayService.PushStream:output_type -> ecsrelay.PushResponse
	7,  // 25: ecsrelay.ECSRelayService.Push:output_type -> ecsrelay.PushResponse
	11, // 26: ecsrelay.ECSRelayService.MinBalanceForPush:output_type -> ecsrelay.BalanceResponse
	16, // [16:27] is the sub-list for method output_type
	5,  // [5:16] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_ecs_relay_proto_init() }
func file_proto_ecs_relay_proto_init() {
	if File_proto_ecs_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ecs_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushManyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountIdentitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountIdentitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ecs_relay_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ecs_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ecs_relay_proto_goTypes,
		DependencyIndexes: file_proto_ecs_relay_proto_depIdxs,
		MessageInfos:      file_proto_ecs_relay_proto_msgTypes,
	}.Build()
	File_proto_ecs_relay_proto = out.File
	file_proto_ecs_relay_proto_rawDesc = nil
	file_proto_ecs_relay_proto_goTypes = nil
	file_proto_ecs_relay_proto_depIdxs = nil
}
