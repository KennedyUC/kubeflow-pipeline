// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: backend/api/v1beta1/resource_reference.proto

package go_client

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

type ResourceType int32

const (
	ResourceType_UNKNOWN_RESOURCE_TYPE ResourceType = 0
	ResourceType_EXPERIMENT            ResourceType = 1
	ResourceType_JOB                   ResourceType = 2
	ResourceType_PIPELINE              ResourceType = 3
	ResourceType_PIPELINE_VERSION      ResourceType = 4
	ResourceType_NAMESPACE             ResourceType = 5
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "UNKNOWN_RESOURCE_TYPE",
		1: "EXPERIMENT",
		2: "JOB",
		3: "PIPELINE",
		4: "PIPELINE_VERSION",
		5: "NAMESPACE",
	}
	ResourceType_value = map[string]int32{
		"UNKNOWN_RESOURCE_TYPE": 0,
		"EXPERIMENT":            1,
		"JOB":                   2,
		"PIPELINE":              3,
		"PIPELINE_VERSION":      4,
		"NAMESPACE":             5,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_api_v1beta1_resource_reference_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_backend_api_v1beta1_resource_reference_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_backend_api_v1beta1_resource_reference_proto_rawDescGZIP(), []int{0}
}

type Relationship int32

const (
	Relationship_UNKNOWN_RELATIONSHIP Relationship = 0
	Relationship_OWNER                Relationship = 1
	Relationship_CREATOR              Relationship = 2
)

// Enum value maps for Relationship.
var (
	Relationship_name = map[int32]string{
		0: "UNKNOWN_RELATIONSHIP",
		1: "OWNER",
		2: "CREATOR",
	}
	Relationship_value = map[string]int32{
		"UNKNOWN_RELATIONSHIP": 0,
		"OWNER":                1,
		"CREATOR":              2,
	}
)

func (x Relationship) Enum() *Relationship {
	p := new(Relationship)
	*p = x
	return p
}

func (x Relationship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Relationship) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_api_v1beta1_resource_reference_proto_enumTypes[1].Descriptor()
}

func (Relationship) Type() protoreflect.EnumType {
	return &file_backend_api_v1beta1_resource_reference_proto_enumTypes[1]
}

func (x Relationship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Relationship.Descriptor instead.
func (Relationship) EnumDescriptor() ([]byte, []int) {
	return file_backend_api_v1beta1_resource_reference_proto_rawDescGZIP(), []int{1}
}

type ResourceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the resource that referred to.
	Type ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=api.ResourceType" json:"type,omitempty"`
	// The ID of the resource that referred to.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResourceKey) Reset() {
	*x = ResourceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v1beta1_resource_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceKey) ProtoMessage() {}

func (x *ResourceKey) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v1beta1_resource_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceKey.ProtoReflect.Descriptor instead.
func (*ResourceKey) Descriptor() ([]byte, []int) {
	return file_backend_api_v1beta1_resource_reference_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceKey) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_UNKNOWN_RESOURCE_TYPE
}

func (x *ResourceKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResourceReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *ResourceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The name of the resource that referred to.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Required field. The relationship from referred resource to the object.
	Relationship Relationship `protobuf:"varint,2,opt,name=relationship,proto3,enum=api.Relationship" json:"relationship,omitempty"`
}

func (x *ResourceReference) Reset() {
	*x = ResourceReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v1beta1_resource_reference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceReference) ProtoMessage() {}

func (x *ResourceReference) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v1beta1_resource_reference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceReference.ProtoReflect.Descriptor instead.
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return file_backend_api_v1beta1_resource_reference_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceReference) GetKey() *ResourceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ResourceReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceReference) GetRelationship() Relationship {
	if x != nil {
		return x.Relationship
	}
	return Relationship_UNKNOWN_RELATIONSHIP
}

var File_backend_api_v1beta1_resource_reference_proto protoreflect.FileDescriptor

var file_backend_api_v1beta1_resource_reference_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0x44, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x22, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2a, 0x75,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50,
	0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x4f, 0x42,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x05, 0x2a, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x6f, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_api_v1beta1_resource_reference_proto_rawDescOnce sync.Once
	file_backend_api_v1beta1_resource_reference_proto_rawDescData = file_backend_api_v1beta1_resource_reference_proto_rawDesc
)

func file_backend_api_v1beta1_resource_reference_proto_rawDescGZIP() []byte {
	file_backend_api_v1beta1_resource_reference_proto_rawDescOnce.Do(func() {
		file_backend_api_v1beta1_resource_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_api_v1beta1_resource_reference_proto_rawDescData)
	})
	return file_backend_api_v1beta1_resource_reference_proto_rawDescData
}

var file_backend_api_v1beta1_resource_reference_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_backend_api_v1beta1_resource_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backend_api_v1beta1_resource_reference_proto_goTypes = []interface{}{
	(ResourceType)(0),         // 0: api.ResourceType
	(Relationship)(0),         // 1: api.Relationship
	(*ResourceKey)(nil),       // 2: api.ResourceKey
	(*ResourceReference)(nil), // 3: api.ResourceReference
}
var file_backend_api_v1beta1_resource_reference_proto_depIdxs = []int32{
	0, // 0: api.ResourceKey.type:type_name -> api.ResourceType
	2, // 1: api.ResourceReference.key:type_name -> api.ResourceKey
	1, // 2: api.ResourceReference.relationship:type_name -> api.Relationship
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_backend_api_v1beta1_resource_reference_proto_init() }
func file_backend_api_v1beta1_resource_reference_proto_init() {
	if File_backend_api_v1beta1_resource_reference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_api_v1beta1_resource_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceKey); i {
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
		file_backend_api_v1beta1_resource_reference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceReference); i {
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
			RawDescriptor: file_backend_api_v1beta1_resource_reference_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_api_v1beta1_resource_reference_proto_goTypes,
		DependencyIndexes: file_backend_api_v1beta1_resource_reference_proto_depIdxs,
		EnumInfos:         file_backend_api_v1beta1_resource_reference_proto_enumTypes,
		MessageInfos:      file_backend_api_v1beta1_resource_reference_proto_msgTypes,
	}.Build()
	File_backend_api_v1beta1_resource_reference_proto = out.File
	file_backend_api_v1beta1_resource_reference_proto_rawDesc = nil
	file_backend_api_v1beta1_resource_reference_proto_goTypes = nil
	file_backend_api_v1beta1_resource_reference_proto_depIdxs = nil
}
