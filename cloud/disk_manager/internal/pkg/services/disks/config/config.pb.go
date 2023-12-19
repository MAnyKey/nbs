// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/internal/pkg/services/disks/config/config.proto

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

type DisksConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultBlockSize                           *uint32  `protobuf:"varint,1,opt,name=DefaultBlockSize,def=4096" json:"DefaultBlockSize,omitempty"`
	DeletedDiskExpirationTimeout               *string  `protobuf:"bytes,2,opt,name=DeletedDiskExpirationTimeout,def=30m" json:"DeletedDiskExpirationTimeout,omitempty"`
	ClearDeletedDisksTaskScheduleInterval      *string  `protobuf:"bytes,3,opt,name=ClearDeletedDisksTaskScheduleInterval,def=1m" json:"ClearDeletedDisksTaskScheduleInterval,omitempty"`
	ClearDeletedDisksLimit                     *uint32  `protobuf:"varint,4,opt,name=ClearDeletedDisksLimit,def=1000" json:"ClearDeletedDisksLimit,omitempty"`
	StorageFolder                              *string  `protobuf:"bytes,5,opt,name=StorageFolder,def=disks" json:"StorageFolder,omitempty"`
	DisableOverlayDisks                        *bool    `protobuf:"varint,6,opt,name=DisableOverlayDisks" json:"DisableOverlayDisks,omitempty"`
	OverlayDisksFolderIdWhitelist              []string `protobuf:"bytes,7,rep,name=OverlayDisksFolderIdWhitelist" json:"OverlayDisksFolderIdWhitelist,omitempty"`
	OverlayDisksFolderIdBlacklist              []string `protobuf:"bytes,8,rep,name=OverlayDisksFolderIdBlacklist" json:"OverlayDisksFolderIdBlacklist,omitempty"`
	EndedMigrationExpirationTimeout            *string  `protobuf:"bytes,9,opt,name=EndedMigrationExpirationTimeout,def=30m" json:"EndedMigrationExpirationTimeout,omitempty"`
	EnableOptimizationForOverlayDiskRelocation *bool    `protobuf:"varint,10,opt,name=EnableOptimizationForOverlayDiskRelocation,def=1" json:"EnableOptimizationForOverlayDiskRelocation,omitempty"`
}

// Default values for DisksConfig fields.
const (
	Default_DisksConfig_DefaultBlockSize                           = uint32(4096)
	Default_DisksConfig_DeletedDiskExpirationTimeout               = string("30m")
	Default_DisksConfig_ClearDeletedDisksTaskScheduleInterval      = string("1m")
	Default_DisksConfig_ClearDeletedDisksLimit                     = uint32(1000)
	Default_DisksConfig_StorageFolder                              = string("disks")
	Default_DisksConfig_EndedMigrationExpirationTimeout            = string("30m")
	Default_DisksConfig_EnableOptimizationForOverlayDiskRelocation = bool(true)
)

func (x *DisksConfig) Reset() {
	*x = DisksConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisksConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisksConfig) ProtoMessage() {}

func (x *DisksConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisksConfig.ProtoReflect.Descriptor instead.
func (*DisksConfig) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *DisksConfig) GetDefaultBlockSize() uint32 {
	if x != nil && x.DefaultBlockSize != nil {
		return *x.DefaultBlockSize
	}
	return Default_DisksConfig_DefaultBlockSize
}

func (x *DisksConfig) GetDeletedDiskExpirationTimeout() string {
	if x != nil && x.DeletedDiskExpirationTimeout != nil {
		return *x.DeletedDiskExpirationTimeout
	}
	return Default_DisksConfig_DeletedDiskExpirationTimeout
}

func (x *DisksConfig) GetClearDeletedDisksTaskScheduleInterval() string {
	if x != nil && x.ClearDeletedDisksTaskScheduleInterval != nil {
		return *x.ClearDeletedDisksTaskScheduleInterval
	}
	return Default_DisksConfig_ClearDeletedDisksTaskScheduleInterval
}

func (x *DisksConfig) GetClearDeletedDisksLimit() uint32 {
	if x != nil && x.ClearDeletedDisksLimit != nil {
		return *x.ClearDeletedDisksLimit
	}
	return Default_DisksConfig_ClearDeletedDisksLimit
}

func (x *DisksConfig) GetStorageFolder() string {
	if x != nil && x.StorageFolder != nil {
		return *x.StorageFolder
	}
	return Default_DisksConfig_StorageFolder
}

func (x *DisksConfig) GetDisableOverlayDisks() bool {
	if x != nil && x.DisableOverlayDisks != nil {
		return *x.DisableOverlayDisks
	}
	return false
}

func (x *DisksConfig) GetOverlayDisksFolderIdWhitelist() []string {
	if x != nil {
		return x.OverlayDisksFolderIdWhitelist
	}
	return nil
}

func (x *DisksConfig) GetOverlayDisksFolderIdBlacklist() []string {
	if x != nil {
		return x.OverlayDisksFolderIdBlacklist
	}
	return nil
}

func (x *DisksConfig) GetEndedMigrationExpirationTimeout() string {
	if x != nil && x.EndedMigrationExpirationTimeout != nil {
		return *x.EndedMigrationExpirationTimeout
	}
	return Default_DisksConfig_EndedMigrationExpirationTimeout
}

func (x *DisksConfig) GetEnableOptimizationForOverlayDiskRelocation() bool {
	if x != nil && x.EnableOptimizationForOverlayDiskRelocation != nil {
		return *x.EnableOptimizationForOverlayDiskRelocation
	}
	return Default_DisksConfig_EnableOptimizationForOverlayDiskRelocation
}

var File_cloud_disk_manager_internal_pkg_services_disks_config_config_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDesc = []byte{
	0x0a, 0x42, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x22, 0xc0, 0x05, 0x0a, 0x0b,
	0x44, 0x69, 0x73, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x10, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x04, 0x34, 0x30, 0x39, 0x36, 0x52, 0x10, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a,
	0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x3a, 0x03, 0x33, 0x30, 0x6d, 0x52, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x44, 0x69, 0x73, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x58, 0x0a, 0x25, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x02, 0x31, 0x6d, 0x52, 0x25, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x3c, 0x0a, 0x16, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x44, 0x69, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x3a, 0x04, 0x31, 0x30, 0x30, 0x30, 0x52, 0x16, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2b,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73,
	0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x44, 0x0a,
	0x1d, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73,
	0x6b, 0x73, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x1d, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69,
	0x73, 0x6b, 0x73, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x1f, 0x45, 0x6e, 0x64,
	0x65, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x3a, 0x03, 0x33, 0x30, 0x6d, 0x52, 0x1f, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x64, 0x0a, 0x2a, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x72, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72,
	0x75, 0x65, 0x52, 0x2a, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x53,
	0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67,
}

var (
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescData = file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDesc
)

func file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescData)
	})
	return file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDescData
}

var file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_goTypes = []interface{}{
	(*DisksConfig)(nil), // 0: disks.DisksConfig
}
var file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_init() }
func file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_init() {
	if File_cloud_disk_manager_internal_pkg_services_disks_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisksConfig); i {
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
			RawDescriptor: file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_depIdxs,
		MessageInfos:      file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_internal_pkg_services_disks_config_config_proto = out.File
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_rawDesc = nil
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_goTypes = nil
	file_cloud_disk_manager_internal_pkg_services_disks_config_config_proto_depIdxs = nil
}
