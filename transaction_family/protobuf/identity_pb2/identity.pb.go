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
// source: protobuf/identity_pb2/identity.proto

package identity

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

type Policy_EntryType int32

const (
	Policy_ENTRY_TYPE_UNSET Policy_EntryType = 0
	Policy_PERMIT_KEY       Policy_EntryType = 1
	Policy_DENY_KEY         Policy_EntryType = 2
)

// Enum value maps for Policy_EntryType.
var (
	Policy_EntryType_name = map[int32]string{
		0: "ENTRY_TYPE_UNSET",
		1: "PERMIT_KEY",
		2: "DENY_KEY",
	}
	Policy_EntryType_value = map[string]int32{
		"ENTRY_TYPE_UNSET": 0,
		"PERMIT_KEY":       1,
		"DENY_KEY":         2,
	}
)

func (x Policy_EntryType) Enum() *Policy_EntryType {
	p := new(Policy_EntryType)
	*p = x
	return p
}

func (x Policy_EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Policy_EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_identity_pb2_identity_proto_enumTypes[0].Descriptor()
}

func (Policy_EntryType) Type() protoreflect.EnumType {
	return &file_protobuf_identity_pb2_identity_proto_enumTypes[0]
}

func (x Policy_EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Policy_EntryType.Descriptor instead.
func (Policy_EntryType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{0, 0}
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the policy, this should be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// list of Entries
	// The entries will be processed in order from first to last.
	Entries []*Policy_Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetEntries() []*Policy_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// Policy will be stored in a Policy list to account for state collisions
type PolicyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *PolicyList) Reset() {
	*x = PolicyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyList) ProtoMessage() {}

func (x *PolicyList) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyList.ProtoReflect.Descriptor instead.
func (*PolicyList) Descriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyList) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of corresponding policy
	PolicyName string `protobuf:"bytes,2,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

// Roles will be stored in a RoleList to account for state collisions
type RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *RoleList) Reset() {
	*x = RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleList) ProtoMessage() {}

func (x *RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleList.ProtoReflect.Descriptor instead.
func (*RoleList) Descriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{3}
}

func (x *RoleList) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Policy_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this is a Permit_KEY or Deny_KEY entry
	Type Policy_EntryType `protobuf:"varint,1,opt,name=type,proto3,enum=Policy_EntryType" json:"type,omitempty"`
	// This should be a list of public keys or * to refer to all participants.
	// If using *, it should be the only key in the list.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Policy_Entry) Reset() {
	*x = Policy_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Entry) ProtoMessage() {}

func (x *Policy_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_identity_pb2_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Entry.ProtoReflect.Descriptor instead.
func (*Policy_Entry) Descriptor() ([]byte, []int) {
	return file_protobuf_identity_pb2_identity_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Policy_Entry) GetType() Policy_EntryType {
	if x != nil {
		return x.Type
	}
	return Policy_ENTRY_TYPE_UNSET
}

func (x *Policy_Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_protobuf_identity_pb2_identity_proto protoreflect.FileDescriptor

var file_protobuf_identity_pb2_identity_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x40,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x3f, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45,
	0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4e, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x10,
	0x02, 0x22, 0x31, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x27, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x1e, 0x0a, 0x1a, 0x73, 0x61,
	0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobuf_identity_pb2_identity_proto_rawDescOnce sync.Once
	file_protobuf_identity_pb2_identity_proto_rawDescData = file_protobuf_identity_pb2_identity_proto_rawDesc
)

func file_protobuf_identity_pb2_identity_proto_rawDescGZIP() []byte {
	file_protobuf_identity_pb2_identity_proto_rawDescOnce.Do(func() {
		file_protobuf_identity_pb2_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_identity_pb2_identity_proto_rawDescData)
	})
	return file_protobuf_identity_pb2_identity_proto_rawDescData
}

var file_protobuf_identity_pb2_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_identity_pb2_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_identity_pb2_identity_proto_goTypes = []interface{}{
	(Policy_EntryType)(0), // 0: Policy.EntryType
	(*Policy)(nil),        // 1: Policy
	(*PolicyList)(nil),    // 2: PolicyList
	(*Role)(nil),          // 3: Role
	(*RoleList)(nil),      // 4: RoleList
	(*Policy_Entry)(nil),  // 5: Policy.Entry
}
var file_protobuf_identity_pb2_identity_proto_depIdxs = []int32{
	5, // 0: Policy.entries:type_name -> Policy.Entry
	1, // 1: PolicyList.policies:type_name -> Policy
	3, // 2: RoleList.roles:type_name -> Role
	0, // 3: Policy.Entry.type:type_name -> Policy.EntryType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_identity_pb2_identity_proto_init() }
func file_protobuf_identity_pb2_identity_proto_init() {
	if File_protobuf_identity_pb2_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_identity_pb2_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_protobuf_identity_pb2_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyList); i {
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
		file_protobuf_identity_pb2_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_protobuf_identity_pb2_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleList); i {
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
		file_protobuf_identity_pb2_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_Entry); i {
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
			RawDescriptor: file_protobuf_identity_pb2_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_identity_pb2_identity_proto_goTypes,
		DependencyIndexes: file_protobuf_identity_pb2_identity_proto_depIdxs,
		EnumInfos:         file_protobuf_identity_pb2_identity_proto_enumTypes,
		MessageInfos:      file_protobuf_identity_pb2_identity_proto_msgTypes,
	}.Build()
	File_protobuf_identity_pb2_identity_proto = out.File
	file_protobuf_identity_pb2_identity_proto_rawDesc = nil
	file_protobuf_identity_pb2_identity_proto_goTypes = nil
	file_protobuf_identity_pb2_identity_proto_depIdxs = nil
}
