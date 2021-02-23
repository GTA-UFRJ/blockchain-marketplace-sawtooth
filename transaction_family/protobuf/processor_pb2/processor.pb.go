// Copyright 2016 Intel Corporation
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
// source: protobuf/processor_pb2/processor.proto

package processor_pb2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	transaction_pb2 "protobuf/transaction_pb2"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TpRegisterResponse_Status int32

const (
	TpRegisterResponse_STATUS_UNSET TpRegisterResponse_Status = 0
	TpRegisterResponse_OK           TpRegisterResponse_Status = 1
	TpRegisterResponse_ERROR        TpRegisterResponse_Status = 2
)

// Enum value maps for TpRegisterResponse_Status.
var (
	TpRegisterResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "ERROR",
	}
	TpRegisterResponse_Status_value = map[string]int32{
		"STATUS_UNSET": 0,
		"OK":           1,
		"ERROR":        2,
	}
)

func (x TpRegisterResponse_Status) Enum() *TpRegisterResponse_Status {
	p := new(TpRegisterResponse_Status)
	*p = x
	return p
}

func (x TpRegisterResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TpRegisterResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_processor_pb2_processor_proto_enumTypes[0].Descriptor()
}

func (TpRegisterResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_processor_pb2_processor_proto_enumTypes[0]
}

func (x TpRegisterResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TpRegisterResponse_Status.Descriptor instead.
func (TpRegisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{1, 0}
}

type TpUnregisterResponse_Status int32

const (
	TpUnregisterResponse_STATUS_UNSET TpUnregisterResponse_Status = 0
	TpUnregisterResponse_OK           TpUnregisterResponse_Status = 1
	TpUnregisterResponse_ERROR        TpUnregisterResponse_Status = 2
)

// Enum value maps for TpUnregisterResponse_Status.
var (
	TpUnregisterResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "ERROR",
	}
	TpUnregisterResponse_Status_value = map[string]int32{
		"STATUS_UNSET": 0,
		"OK":           1,
		"ERROR":        2,
	}
)

func (x TpUnregisterResponse_Status) Enum() *TpUnregisterResponse_Status {
	p := new(TpUnregisterResponse_Status)
	*p = x
	return p
}

func (x TpUnregisterResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TpUnregisterResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_processor_pb2_processor_proto_enumTypes[1].Descriptor()
}

func (TpUnregisterResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_processor_pb2_processor_proto_enumTypes[1]
}

func (x TpUnregisterResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TpUnregisterResponse_Status.Descriptor instead.
func (TpUnregisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{3, 0}
}

type TpProcessResponse_Status int32

const (
	TpProcessResponse_STATUS_UNSET        TpProcessResponse_Status = 0
	TpProcessResponse_OK                  TpProcessResponse_Status = 1
	TpProcessResponse_INVALID_TRANSACTION TpProcessResponse_Status = 2
	TpProcessResponse_INTERNAL_ERROR      TpProcessResponse_Status = 3
)

// Enum value maps for TpProcessResponse_Status.
var (
	TpProcessResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INVALID_TRANSACTION",
		3: "INTERNAL_ERROR",
	}
	TpProcessResponse_Status_value = map[string]int32{
		"STATUS_UNSET":        0,
		"OK":                  1,
		"INVALID_TRANSACTION": 2,
		"INTERNAL_ERROR":      3,
	}
)

func (x TpProcessResponse_Status) Enum() *TpProcessResponse_Status {
	p := new(TpProcessResponse_Status)
	*p = x
	return p
}

func (x TpProcessResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TpProcessResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_processor_pb2_processor_proto_enumTypes[2].Descriptor()
}

func (TpProcessResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_processor_pb2_processor_proto_enumTypes[2]
}

func (x TpProcessResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TpProcessResponse_Status.Descriptor instead.
func (TpProcessResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{5, 0}
}

// The registration request from the transaction processor to the
// validator/executor
type TpRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A settled upon name for the capabilities of the transaction processor.
	// For example: intkey, xo
	Family string `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	// The version supported.  For example:
	//      1.0  for version 1.0
	//      2.1  for version 2.1
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The namespaces this transaction processor expects to interact with
	// when processing transactions matching this specification; will be
	// enforced by the state API on the validator.
	Namespaces []string `protobuf:"bytes,4,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// The maximum number of transactions that this transaction processor can
	// handle at once.
	MaxOccupancy uint32 `protobuf:"varint,5,opt,name=max_occupancy,json=maxOccupancy,proto3" json:"max_occupancy,omitempty"`
}

func (x *TpRegisterRequest) Reset() {
	*x = TpRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpRegisterRequest) ProtoMessage() {}

func (x *TpRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpRegisterRequest.ProtoReflect.Descriptor instead.
func (*TpRegisterRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{0}
}

func (x *TpRegisterRequest) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *TpRegisterRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TpRegisterRequest) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *TpRegisterRequest) GetMaxOccupancy() uint32 {
	if x != nil {
		return x.MaxOccupancy
	}
	return 0
}

// A response sent from the validator to the transaction processor
// acknowledging the registration
type TpRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TpRegisterResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TpRegisterResponse_Status" json:"status,omitempty"`
}

func (x *TpRegisterResponse) Reset() {
	*x = TpRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpRegisterResponse) ProtoMessage() {}

func (x *TpRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpRegisterResponse.ProtoReflect.Descriptor instead.
func (*TpRegisterResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{1}
}

func (x *TpRegisterResponse) GetStatus() TpRegisterResponse_Status {
	if x != nil {
		return x.Status
	}
	return TpRegisterResponse_STATUS_UNSET
}

// The unregistration request from the transaction processor to the
// validator/executor. The correct handlers are determined from the
// zeromq identity of the tp, on the validator side.
type TpUnregisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TpUnregisterRequest) Reset() {
	*x = TpUnregisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpUnregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpUnregisterRequest) ProtoMessage() {}

func (x *TpUnregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpUnregisterRequest.ProtoReflect.Descriptor instead.
func (*TpUnregisterRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{2}
}

// A response sent from the validator to the transaction processor
// acknowledging the unregistration
type TpUnregisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TpUnregisterResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TpUnregisterResponse_Status" json:"status,omitempty"`
}

func (x *TpUnregisterResponse) Reset() {
	*x = TpUnregisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpUnregisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpUnregisterResponse) ProtoMessage() {}

func (x *TpUnregisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpUnregisterResponse.ProtoReflect.Descriptor instead.
func (*TpUnregisterResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{3}
}

func (x *TpUnregisterResponse) GetStatus() TpUnregisterResponse_Status {
	if x != nil {
		return x.Status
	}
	return TpUnregisterResponse_STATUS_UNSET
}

// The request from the validator/executor of the transaction processor
// to verify a transaction.
type TpProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *transaction_pb2.TransactionHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`                        // The transaction header
	Payload   []byte                             `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                      // The transaction payload
	Signature string                             `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`                  // The transaction header_signature
	ContextId string                             `protobuf:"bytes,4,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"` // The context_id for state requests.
}

func (x *TpProcessRequest) Reset() {
	*x = TpProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpProcessRequest) ProtoMessage() {}

func (x *TpProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpProcessRequest.ProtoReflect.Descriptor instead.
func (*TpProcessRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{4}
}

func (x *TpProcessRequest) GetHeader() *transaction_pb2.TransactionHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TpProcessRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TpProcessRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *TpProcessRequest) GetContextId() string {
	if x != nil {
		return x.ContextId
	}
	return ""
}

// The response from the transaction processor to the validator/executor
// used to respond about the validity of a transaction
type TpProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TpProcessResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TpProcessResponse_Status" json:"status,omitempty"`
	// A message to include on responses in the cases where
	// status is either INVALID_TRANSACTION or INTERNAL_ERROR
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Information that may be included with the response.
	// This information is an opaque, application-specific encoded block of
	// data that will be propagated back to the transaction submitter.
	ExtendedData []byte `protobuf:"bytes,3,opt,name=extended_data,json=extendedData,proto3" json:"extended_data,omitempty"`
}

func (x *TpProcessResponse) Reset() {
	*x = TpProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpProcessResponse) ProtoMessage() {}

func (x *TpProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_processor_pb2_processor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpProcessResponse.ProtoReflect.Descriptor instead.
func (*TpProcessResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_processor_pb2_processor_proto_rawDescGZIP(), []int{5}
}

func (x *TpProcessResponse) GetStatus() TpProcessResponse_Status {
	if x != nil {
		return x.Status
	}
	return TpProcessResponse_STATUS_UNSET
}

func (x *TpProcessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TpProcessResponse) GetExtendedData() []byte {
	if x != nil {
		return x.ExtendedData
	}
	return nil
}

var File_protobuf_processor_pb2_processor_proto protoreflect.FileDescriptor

var file_protobuf_processor_pb2_processor_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x62, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x54, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63,
	0x79, 0x22, 0x77, 0x0a, 0x12, 0x54, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x54, 0x70, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x15, 0x0a, 0x13, 0x54, 0x70,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x7b, 0x0a, 0x14, 0x54, 0x70, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x54, 0x70, 0x55, 0x6e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x95,
	0x01, 0x0a, 0x10, 0x54, 0x70, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x11, 0x54, 0x70, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x54,
	0x70, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4f,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42,
	0x28, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_processor_pb2_processor_proto_rawDescOnce sync.Once
	file_protobuf_processor_pb2_processor_proto_rawDescData = file_protobuf_processor_pb2_processor_proto_rawDesc
)

func file_protobuf_processor_pb2_processor_proto_rawDescGZIP() []byte {
	file_protobuf_processor_pb2_processor_proto_rawDescOnce.Do(func() {
		file_protobuf_processor_pb2_processor_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_processor_pb2_processor_proto_rawDescData)
	})
	return file_protobuf_processor_pb2_processor_proto_rawDescData
}

var file_protobuf_processor_pb2_processor_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protobuf_processor_pb2_processor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_processor_pb2_processor_proto_goTypes = []interface{}{
	(TpRegisterResponse_Status)(0),            // 0: TpRegisterResponse.Status
	(TpUnregisterResponse_Status)(0),          // 1: TpUnregisterResponse.Status
	(TpProcessResponse_Status)(0),             // 2: TpProcessResponse.Status
	(*TpRegisterRequest)(nil),                 // 3: TpRegisterRequest
	(*TpRegisterResponse)(nil),                // 4: TpRegisterResponse
	(*TpUnregisterRequest)(nil),               // 5: TpUnregisterRequest
	(*TpUnregisterResponse)(nil),              // 6: TpUnregisterResponse
	(*TpProcessRequest)(nil),                  // 7: TpProcessRequest
	(*TpProcessResponse)(nil),                 // 8: TpProcessResponse
	(*transaction_pb2.TransactionHeader)(nil), // 9: TransactionHeader
}
var file_protobuf_processor_pb2_processor_proto_depIdxs = []int32{
	0, // 0: TpRegisterResponse.status:type_name -> TpRegisterResponse.Status
	1, // 1: TpUnregisterResponse.status:type_name -> TpUnregisterResponse.Status
	9, // 2: TpProcessRequest.header:type_name -> TransactionHeader
	2, // 3: TpProcessResponse.status:type_name -> TpProcessResponse.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_processor_pb2_processor_proto_init() }
func file_protobuf_processor_pb2_processor_proto_init() {
	if File_protobuf_processor_pb2_processor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_processor_pb2_processor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpRegisterRequest); i {
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
		file_protobuf_processor_pb2_processor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpRegisterResponse); i {
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
		file_protobuf_processor_pb2_processor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpUnregisterRequest); i {
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
		file_protobuf_processor_pb2_processor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpUnregisterResponse); i {
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
		file_protobuf_processor_pb2_processor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpProcessRequest); i {
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
		file_protobuf_processor_pb2_processor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpProcessResponse); i {
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
			RawDescriptor: file_protobuf_processor_pb2_processor_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_processor_pb2_processor_proto_goTypes,
		DependencyIndexes: file_protobuf_processor_pb2_processor_proto_depIdxs,
		EnumInfos:         file_protobuf_processor_pb2_processor_proto_enumTypes,
		MessageInfos:      file_protobuf_processor_pb2_processor_proto_msgTypes,
	}.Build()
	File_protobuf_processor_pb2_processor_proto = out.File
	file_protobuf_processor_pb2_processor_proto_rawDesc = nil
	file_protobuf_processor_pb2_processor_proto_goTypes = nil
	file_protobuf_processor_pb2_processor_proto_depIdxs = nil
}
