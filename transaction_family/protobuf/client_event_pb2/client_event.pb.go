// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: protobuf/client_event_pb2/client_event.proto

package client_event_pb2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	events_pb2 "protobuf/events_pb2"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientEventsSubscribeResponse_Status int32

const (
	ClientEventsSubscribeResponse_STATUS_UNSET   ClientEventsSubscribeResponse_Status = 0
	ClientEventsSubscribeResponse_OK             ClientEventsSubscribeResponse_Status = 1
	ClientEventsSubscribeResponse_INVALID_FILTER ClientEventsSubscribeResponse_Status = 2
	ClientEventsSubscribeResponse_UNKNOWN_BLOCK  ClientEventsSubscribeResponse_Status = 3
)

// Enum value maps for ClientEventsSubscribeResponse_Status.
var (
	ClientEventsSubscribeResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INVALID_FILTER",
		3: "UNKNOWN_BLOCK",
	}
	ClientEventsSubscribeResponse_Status_value = map[string]int32{
		"STATUS_UNSET":   0,
		"OK":             1,
		"INVALID_FILTER": 2,
		"UNKNOWN_BLOCK":  3,
	}
)

func (x ClientEventsSubscribeResponse_Status) Enum() *ClientEventsSubscribeResponse_Status {
	p := new(ClientEventsSubscribeResponse_Status)
	*p = x
	return p
}

func (x ClientEventsSubscribeResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientEventsSubscribeResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_client_event_pb2_client_event_proto_enumTypes[0].Descriptor()
}

func (ClientEventsSubscribeResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_client_event_pb2_client_event_proto_enumTypes[0]
}

func (x ClientEventsSubscribeResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientEventsSubscribeResponse_Status.Descriptor instead.
func (ClientEventsSubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{1, 0}
}

type ClientEventsUnsubscribeResponse_Status int32

const (
	ClientEventsUnsubscribeResponse_STATUS_UNSET   ClientEventsUnsubscribeResponse_Status = 0
	ClientEventsUnsubscribeResponse_OK             ClientEventsUnsubscribeResponse_Status = 1
	ClientEventsUnsubscribeResponse_INTERNAL_ERROR ClientEventsUnsubscribeResponse_Status = 2
)

// Enum value maps for ClientEventsUnsubscribeResponse_Status.
var (
	ClientEventsUnsubscribeResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INTERNAL_ERROR",
	}
	ClientEventsUnsubscribeResponse_Status_value = map[string]int32{
		"STATUS_UNSET":   0,
		"OK":             1,
		"INTERNAL_ERROR": 2,
	}
)

func (x ClientEventsUnsubscribeResponse_Status) Enum() *ClientEventsUnsubscribeResponse_Status {
	p := new(ClientEventsUnsubscribeResponse_Status)
	*p = x
	return p
}

func (x ClientEventsUnsubscribeResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientEventsUnsubscribeResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_client_event_pb2_client_event_proto_enumTypes[1].Descriptor()
}

func (ClientEventsUnsubscribeResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_client_event_pb2_client_event_proto_enumTypes[1]
}

func (x ClientEventsUnsubscribeResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientEventsUnsubscribeResponse_Status.Descriptor instead.
func (ClientEventsUnsubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{3, 0}
}

type ClientEventsGetResponse_Status int32

const (
	ClientEventsGetResponse_STATUS_UNSET   ClientEventsGetResponse_Status = 0
	ClientEventsGetResponse_OK             ClientEventsGetResponse_Status = 1
	ClientEventsGetResponse_INTERNAL_ERROR ClientEventsGetResponse_Status = 2
	ClientEventsGetResponse_INVALID_FILTER ClientEventsGetResponse_Status = 3
	ClientEventsGetResponse_UNKNOWN_BLOCK  ClientEventsGetResponse_Status = 4
)

// Enum value maps for ClientEventsGetResponse_Status.
var (
	ClientEventsGetResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INTERNAL_ERROR",
		3: "INVALID_FILTER",
		4: "UNKNOWN_BLOCK",
	}
	ClientEventsGetResponse_Status_value = map[string]int32{
		"STATUS_UNSET":   0,
		"OK":             1,
		"INTERNAL_ERROR": 2,
		"INVALID_FILTER": 3,
		"UNKNOWN_BLOCK":  4,
	}
)

func (x ClientEventsGetResponse_Status) Enum() *ClientEventsGetResponse_Status {
	p := new(ClientEventsGetResponse_Status)
	*p = x
	return p
}

func (x ClientEventsGetResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientEventsGetResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_client_event_pb2_client_event_proto_enumTypes[2].Descriptor()
}

func (ClientEventsGetResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_client_event_pb2_client_event_proto_enumTypes[2]
}

func (x ClientEventsGetResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientEventsGetResponse_Status.Descriptor instead.
func (ClientEventsGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{5, 0}
}

type ClientEventsSubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*events_pb2.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	// The block id (or ids, if trying to walk back a fork) the subscriber last
	// received events on. It can be set to empty if it has not yet received the
	// genesis block.
	LastKnownBlockIds []string `protobuf:"bytes,2,rep,name=last_known_block_ids,json=lastKnownBlockIds,proto3" json:"last_known_block_ids,omitempty"`
}

func (x *ClientEventsSubscribeRequest) Reset() {
	*x = ClientEventsSubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsSubscribeRequest) ProtoMessage() {}

func (x *ClientEventsSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsSubscribeRequest.ProtoReflect.Descriptor instead.
func (*ClientEventsSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{0}
}

func (x *ClientEventsSubscribeRequest) GetSubscriptions() []*events_pb2.EventSubscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *ClientEventsSubscribeRequest) GetLastKnownBlockIds() []string {
	if x != nil {
		return x.LastKnownBlockIds
	}
	return nil
}

type ClientEventsSubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ClientEventsSubscribeResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ClientEventsSubscribeResponse_Status" json:"status,omitempty"`
	// Additional information about the response status
	ResponseMessage string `protobuf:"bytes,2,opt,name=response_message,json=responseMessage,proto3" json:"response_message,omitempty"`
}

func (x *ClientEventsSubscribeResponse) Reset() {
	*x = ClientEventsSubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsSubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsSubscribeResponse) ProtoMessage() {}

func (x *ClientEventsSubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsSubscribeResponse.ProtoReflect.Descriptor instead.
func (*ClientEventsSubscribeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{1}
}

func (x *ClientEventsSubscribeResponse) GetStatus() ClientEventsSubscribeResponse_Status {
	if x != nil {
		return x.Status
	}
	return ClientEventsSubscribeResponse_STATUS_UNSET
}

func (x *ClientEventsSubscribeResponse) GetResponseMessage() string {
	if x != nil {
		return x.ResponseMessage
	}
	return ""
}

type ClientEventsUnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientEventsUnsubscribeRequest) Reset() {
	*x = ClientEventsUnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsUnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsUnsubscribeRequest) ProtoMessage() {}

func (x *ClientEventsUnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsUnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*ClientEventsUnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{2}
}

type ClientEventsUnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ClientEventsUnsubscribeResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ClientEventsUnsubscribeResponse_Status" json:"status,omitempty"`
}

func (x *ClientEventsUnsubscribeResponse) Reset() {
	*x = ClientEventsUnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsUnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsUnsubscribeResponse) ProtoMessage() {}

func (x *ClientEventsUnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsUnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*ClientEventsUnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{3}
}

func (x *ClientEventsUnsubscribeResponse) GetStatus() ClientEventsUnsubscribeResponse_Status {
	if x != nil {
		return x.Status
	}
	return ClientEventsUnsubscribeResponse_STATUS_UNSET
}

type ClientEventsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*events_pb2.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	BlockIds      []string                        `protobuf:"bytes,2,rep,name=block_ids,json=blockIds,proto3" json:"block_ids,omitempty"`
}

func (x *ClientEventsGetRequest) Reset() {
	*x = ClientEventsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsGetRequest) ProtoMessage() {}

func (x *ClientEventsGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsGetRequest.ProtoReflect.Descriptor instead.
func (*ClientEventsGetRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{4}
}

func (x *ClientEventsGetRequest) GetSubscriptions() []*events_pb2.EventSubscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *ClientEventsGetRequest) GetBlockIds() []string {
	if x != nil {
		return x.BlockIds
	}
	return nil
}

type ClientEventsGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ClientEventsGetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ClientEventsGetResponse_Status" json:"status,omitempty"`
	Events []*events_pb2.Event            `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ClientEventsGetResponse) Reset() {
	*x = ClientEventsGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventsGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventsGetResponse) ProtoMessage() {}

func (x *ClientEventsGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_event_pb2_client_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventsGetResponse.ProtoReflect.Descriptor instead.
func (*ClientEventsGetResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP(), []int{5}
}

func (x *ClientEventsGetResponse) GetStatus() ClientEventsGetResponse_Status {
	if x != nil {
		return x.Status
	}
	return ClientEventsGetResponse_STATUS_UNSET
}

func (x *ClientEventsGetResponse) GetEvents() []*events_pb2.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_protobuf_client_event_pb2_client_event_proto protoreflect.FileDescriptor

var file_protobuf_client_event_pb2_client_event_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x70, 0x62, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x01, 0x0a, 0x1c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x73, 0x22, 0xd4, 0x01, 0x0a,
	0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x10, 0x03, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x1f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x22, 0x6f, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x04, 0x42, 0x2b, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f,
	0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x50, 0x01, 0x5a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_client_event_pb2_client_event_proto_rawDescOnce sync.Once
	file_protobuf_client_event_pb2_client_event_proto_rawDescData = file_protobuf_client_event_pb2_client_event_proto_rawDesc
)

func file_protobuf_client_event_pb2_client_event_proto_rawDescGZIP() []byte {
	file_protobuf_client_event_pb2_client_event_proto_rawDescOnce.Do(func() {
		file_protobuf_client_event_pb2_client_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_client_event_pb2_client_event_proto_rawDescData)
	})
	return file_protobuf_client_event_pb2_client_event_proto_rawDescData
}

var file_protobuf_client_event_pb2_client_event_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protobuf_client_event_pb2_client_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_client_event_pb2_client_event_proto_goTypes = []interface{}{
	(ClientEventsSubscribeResponse_Status)(0),   // 0: ClientEventsSubscribeResponse.Status
	(ClientEventsUnsubscribeResponse_Status)(0), // 1: ClientEventsUnsubscribeResponse.Status
	(ClientEventsGetResponse_Status)(0),         // 2: ClientEventsGetResponse.Status
	(*ClientEventsSubscribeRequest)(nil),        // 3: ClientEventsSubscribeRequest
	(*ClientEventsSubscribeResponse)(nil),       // 4: ClientEventsSubscribeResponse
	(*ClientEventsUnsubscribeRequest)(nil),      // 5: ClientEventsUnsubscribeRequest
	(*ClientEventsUnsubscribeResponse)(nil),     // 6: ClientEventsUnsubscribeResponse
	(*ClientEventsGetRequest)(nil),              // 7: ClientEventsGetRequest
	(*ClientEventsGetResponse)(nil),             // 8: ClientEventsGetResponse
	(*events_pb2.EventSubscription)(nil),        // 9: EventSubscription
	(*events_pb2.Event)(nil),                    // 10: Event
}
var file_protobuf_client_event_pb2_client_event_proto_depIdxs = []int32{
	9,  // 0: ClientEventsSubscribeRequest.subscriptions:type_name -> EventSubscription
	0,  // 1: ClientEventsSubscribeResponse.status:type_name -> ClientEventsSubscribeResponse.Status
	1,  // 2: ClientEventsUnsubscribeResponse.status:type_name -> ClientEventsUnsubscribeResponse.Status
	9,  // 3: ClientEventsGetRequest.subscriptions:type_name -> EventSubscription
	2,  // 4: ClientEventsGetResponse.status:type_name -> ClientEventsGetResponse.Status
	10, // 5: ClientEventsGetResponse.events:type_name -> Event
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_protobuf_client_event_pb2_client_event_proto_init() }
func file_protobuf_client_event_pb2_client_event_proto_init() {
	if File_protobuf_client_event_pb2_client_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsSubscribeRequest); i {
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
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsSubscribeResponse); i {
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
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsUnsubscribeRequest); i {
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
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsUnsubscribeResponse); i {
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
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsGetRequest); i {
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
		file_protobuf_client_event_pb2_client_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventsGetResponse); i {
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
			RawDescriptor: file_protobuf_client_event_pb2_client_event_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_client_event_pb2_client_event_proto_goTypes,
		DependencyIndexes: file_protobuf_client_event_pb2_client_event_proto_depIdxs,
		EnumInfos:         file_protobuf_client_event_pb2_client_event_proto_enumTypes,
		MessageInfos:      file_protobuf_client_event_pb2_client_event_proto_msgTypes,
	}.Build()
	File_protobuf_client_event_pb2_client_event_proto = out.File
	file_protobuf_client_event_pb2_client_event_proto_rawDesc = nil
	file_protobuf_client_event_pb2_client_event_proto_goTypes = nil
	file_protobuf_client_event_pb2_client_event_proto_depIdxs = nil
}
