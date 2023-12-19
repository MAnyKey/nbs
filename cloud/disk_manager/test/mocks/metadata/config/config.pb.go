// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/test/mocks/metadata/config/config.proto

package config

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

type MetadataServiceMockConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port        *uint32 `protobuf:"varint,1,req,name=Port" json:"Port,omitempty"`
	AccessToken *string `protobuf:"bytes,2,opt,name=AccessToken" json:"AccessToken,omitempty"`
}

func (x *MetadataServiceMockConfig) Reset() {
	*x = MetadataServiceMockConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_test_mocks_metadata_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataServiceMockConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataServiceMockConfig) ProtoMessage() {}

func (x *MetadataServiceMockConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_test_mocks_metadata_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataServiceMockConfig.ProtoReflect.Descriptor instead.
func (*MetadataServiceMockConfig) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataServiceMockConfig) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *MetadataServiceMockConfig) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

var File_cloud_disk_manager_test_mocks_metadata_config_config_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x51, 0x0a, 0x19, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73,
	0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d,
	0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67,
}

var (
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescData = file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDesc
)

func file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescData)
	})
	return file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDescData
}

var file_cloud_disk_manager_test_mocks_metadata_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_disk_manager_test_mocks_metadata_config_config_proto_goTypes = []interface{}{
	(*MetadataServiceMockConfig)(nil), // 0: config.MetadataServiceMockConfig
}
var file_cloud_disk_manager_test_mocks_metadata_config_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_disk_manager_test_mocks_metadata_config_config_proto_init() }
func file_cloud_disk_manager_test_mocks_metadata_config_config_proto_init() {
	if File_cloud_disk_manager_test_mocks_metadata_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_test_mocks_metadata_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataServiceMockConfig); i {
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
			RawDescriptor: file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_test_mocks_metadata_config_config_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_test_mocks_metadata_config_config_proto_depIdxs,
		MessageInfos:      file_cloud_disk_manager_test_mocks_metadata_config_config_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_test_mocks_metadata_config_config_proto = out.File
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_rawDesc = nil
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_goTypes = nil
	file_cloud_disk_manager_test_mocks_metadata_config_config_proto_depIdxs = nil
}
