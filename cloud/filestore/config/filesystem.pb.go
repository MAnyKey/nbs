// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/filestore/config/filesystem.proto

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

type TFileSystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FileSystem identifier.
	FileSystemId *string `protobuf:"bytes,1,opt,name=FileSystemId" json:"FileSystemId,omitempty"`
	// Filesystem blocksize.
	BlockSize *uint32 `protobuf:"varint,2,opt,name=BlockSize" json:"BlockSize,omitempty"`
	// Keep attempts to acquire lock.
	LockRetryTimeout *uint32 `protobuf:"varint,3,opt,name=LockRetryTimeout" json:"LockRetryTimeout,omitempty"`
	// Inode entry timeout.
	EntryTimeout *uint32 `protobuf:"varint,4,opt,name=EntryTimeout" json:"EntryTimeout,omitempty"`
	// Node attributes timeout.
	AttrTimeout *uint32 `protobuf:"varint,5,opt,name=AttrTimeout" json:"AttrTimeout,omitempty"`
	// X Attrs cache limits
	XAttrCacheLimit   *uint32 `protobuf:"varint,6,opt,name=XAttrCacheLimit" json:"XAttrCacheLimit,omitempty"`
	XAttrCacheTimeout *uint32 `protobuf:"varint,7,opt,name=XAttrCacheTimeout" json:"XAttrCacheTimeout,omitempty"`
	// Filesystem max buffer size per request.
	MaxBufferSize *uint32 `protobuf:"varint,8,opt,name=MaxBufferSize" json:"MaxBufferSize,omitempty"`
}

func (x *TFileSystemConfig) Reset() {
	*x = TFileSystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFileSystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFileSystemConfig) ProtoMessage() {}

func (x *TFileSystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFileSystemConfig.ProtoReflect.Descriptor instead.
func (*TFileSystemConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *TFileSystemConfig) GetFileSystemId() string {
	if x != nil && x.FileSystemId != nil {
		return *x.FileSystemId
	}
	return ""
}

func (x *TFileSystemConfig) GetBlockSize() uint32 {
	if x != nil && x.BlockSize != nil {
		return *x.BlockSize
	}
	return 0
}

func (x *TFileSystemConfig) GetLockRetryTimeout() uint32 {
	if x != nil && x.LockRetryTimeout != nil {
		return *x.LockRetryTimeout
	}
	return 0
}

func (x *TFileSystemConfig) GetEntryTimeout() uint32 {
	if x != nil && x.EntryTimeout != nil {
		return *x.EntryTimeout
	}
	return 0
}

func (x *TFileSystemConfig) GetAttrTimeout() uint32 {
	if x != nil && x.AttrTimeout != nil {
		return *x.AttrTimeout
	}
	return 0
}

func (x *TFileSystemConfig) GetXAttrCacheLimit() uint32 {
	if x != nil && x.XAttrCacheLimit != nil {
		return *x.XAttrCacheLimit
	}
	return 0
}

func (x *TFileSystemConfig) GetXAttrCacheTimeout() uint32 {
	if x != nil && x.XAttrCacheTimeout != nil {
		return *x.XAttrCacheTimeout
	}
	return 0
}

func (x *TFileSystemConfig) GetMaxBufferSize() uint32 {
	if x != nil && x.MaxBufferSize != nil {
		return *x.MaxBufferSize
	}
	return 0
}

var File_cloud_filestore_config_filesystem_proto protoreflect.FileDescriptor

var file_cloud_filestore_config_filesystem_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x4e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x11, 0x54, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x74, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x58, 0x41, 0x74, 0x74, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x58, 0x41, 0x74, 0x74, 0x72, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x58, 0x41, 0x74, 0x74, 0x72,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x58, 0x41, 0x74, 0x74, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x4d, 0x61,
	0x78, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67,
}

var (
	file_cloud_filestore_config_filesystem_proto_rawDescOnce sync.Once
	file_cloud_filestore_config_filesystem_proto_rawDescData = file_cloud_filestore_config_filesystem_proto_rawDesc
)

func file_cloud_filestore_config_filesystem_proto_rawDescGZIP() []byte {
	file_cloud_filestore_config_filesystem_proto_rawDescOnce.Do(func() {
		file_cloud_filestore_config_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_filestore_config_filesystem_proto_rawDescData)
	})
	return file_cloud_filestore_config_filesystem_proto_rawDescData
}

var file_cloud_filestore_config_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_filestore_config_filesystem_proto_goTypes = []interface{}{
	(*TFileSystemConfig)(nil), // 0: NCloud.NFileStore.NProto.TFileSystemConfig
}
var file_cloud_filestore_config_filesystem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_filestore_config_filesystem_proto_init() }
func file_cloud_filestore_config_filesystem_proto_init() {
	if File_cloud_filestore_config_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_filestore_config_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFileSystemConfig); i {
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
			RawDescriptor: file_cloud_filestore_config_filesystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_filestore_config_filesystem_proto_goTypes,
		DependencyIndexes: file_cloud_filestore_config_filesystem_proto_depIdxs,
		MessageInfos:      file_cloud_filestore_config_filesystem_proto_msgTypes,
	}.Build()
	File_cloud_filestore_config_filesystem_proto = out.File
	file_cloud_filestore_config_filesystem_proto_rawDesc = nil
	file_cloud_filestore_config_filesystem_proto_goTypes = nil
	file_cloud_filestore_config_filesystem_proto_depIdxs = nil
}
