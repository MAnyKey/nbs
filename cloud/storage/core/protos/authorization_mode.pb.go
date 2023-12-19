// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/storage/core/protos/authorization_mode.proto

package protos

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

type EAuthorizationMode int32

const (
	// All requests are allowed. Authorization failures are not reported.
	EAuthorizationMode_AUTHORIZATION_IGNORE EAuthorizationMode = 0
	// Requests when authorizer is disabled and requests with empty token are
	// allowed but reported. Other authorization failures fail requests.
	EAuthorizationMode_AUTHORIZATION_ACCEPT EAuthorizationMode = 1
	// All authorization failures fail requests.
	EAuthorizationMode_AUTHORIZATION_REQUIRE EAuthorizationMode = 2
)

// Enum value maps for EAuthorizationMode.
var (
	EAuthorizationMode_name = map[int32]string{
		0: "AUTHORIZATION_IGNORE",
		1: "AUTHORIZATION_ACCEPT",
		2: "AUTHORIZATION_REQUIRE",
	}
	EAuthorizationMode_value = map[string]int32{
		"AUTHORIZATION_IGNORE":  0,
		"AUTHORIZATION_ACCEPT":  1,
		"AUTHORIZATION_REQUIRE": 2,
	}
)

func (x EAuthorizationMode) Enum() *EAuthorizationMode {
	p := new(EAuthorizationMode)
	*p = x
	return p
}

func (x EAuthorizationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EAuthorizationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_storage_core_protos_authorization_mode_proto_enumTypes[0].Descriptor()
}

func (EAuthorizationMode) Type() protoreflect.EnumType {
	return &file_cloud_storage_core_protos_authorization_mode_proto_enumTypes[0]
}

func (x EAuthorizationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EAuthorizationMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EAuthorizationMode(num)
	return nil
}

// Deprecated: Use EAuthorizationMode.Descriptor instead.
func (EAuthorizationMode) EnumDescriptor() ([]byte, []int) {
	return file_cloud_storage_core_protos_authorization_mode_proto_rawDescGZIP(), []int{0}
}

var File_cloud_storage_core_protos_authorization_mode_proto protoreflect.FileDescriptor

var file_cloud_storage_core_protos_authorization_mode_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x63, 0x0a, 0x12, 0x45, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x47, 0x4e, 0x4f, 0x52,
	0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x10, 0x02, 0x42, 0x56, 0x0a, 0x1d, 0x72, 0x75, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_cloud_storage_core_protos_authorization_mode_proto_rawDescOnce sync.Once
	file_cloud_storage_core_protos_authorization_mode_proto_rawDescData = file_cloud_storage_core_protos_authorization_mode_proto_rawDesc
)

func file_cloud_storage_core_protos_authorization_mode_proto_rawDescGZIP() []byte {
	file_cloud_storage_core_protos_authorization_mode_proto_rawDescOnce.Do(func() {
		file_cloud_storage_core_protos_authorization_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_storage_core_protos_authorization_mode_proto_rawDescData)
	})
	return file_cloud_storage_core_protos_authorization_mode_proto_rawDescData
}

var file_cloud_storage_core_protos_authorization_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_storage_core_protos_authorization_mode_proto_goTypes = []interface{}{
	(EAuthorizationMode)(0), // 0: NCloud.NProto.EAuthorizationMode
}
var file_cloud_storage_core_protos_authorization_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_storage_core_protos_authorization_mode_proto_init() }
func file_cloud_storage_core_protos_authorization_mode_proto_init() {
	if File_cloud_storage_core_protos_authorization_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_storage_core_protos_authorization_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_storage_core_protos_authorization_mode_proto_goTypes,
		DependencyIndexes: file_cloud_storage_core_protos_authorization_mode_proto_depIdxs,
		EnumInfos:         file_cloud_storage_core_protos_authorization_mode_proto_enumTypes,
	}.Build()
	File_cloud_storage_core_protos_authorization_mode_proto = out.File
	file_cloud_storage_core_protos_authorization_mode_proto_rawDesc = nil
	file_cloud_storage_core_protos_authorization_mode_proto_goTypes = nil
	file_cloud_storage_core_protos_authorization_mode_proto_depIdxs = nil
}
