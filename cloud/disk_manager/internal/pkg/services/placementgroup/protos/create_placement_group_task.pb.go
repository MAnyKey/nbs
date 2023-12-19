// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/internal/pkg/services/placementgroup/protos/create_placement_group_task.proto

package protos

import (
	types "github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/types"
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

type CreatePlacementGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId                  string                  `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	GroupId                 string                  `protobuf:"bytes,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	PlacementStrategy       types.PlacementStrategy `protobuf:"varint,3,opt,name=PlacementStrategy,proto3,enum=types.PlacementStrategy" json:"PlacementStrategy,omitempty"`
	PlacementPartitionCount uint32                  `protobuf:"varint,5,opt,name=PlacementPartitionCount,proto3" json:"PlacementPartitionCount,omitempty"`
}

func (x *CreatePlacementGroupRequest) Reset() {
	*x = CreatePlacementGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlacementGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlacementGroupRequest) ProtoMessage() {}

func (x *CreatePlacementGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlacementGroupRequest.ProtoReflect.Descriptor instead.
func (*CreatePlacementGroupRequest) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlacementGroupRequest) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *CreatePlacementGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreatePlacementGroupRequest) GetPlacementStrategy() types.PlacementStrategy {
	if x != nil {
		return x.PlacementStrategy
	}
	return types.PlacementStrategy(0)
}

func (x *CreatePlacementGroupRequest) GetPlacementPartitionCount() uint32 {
	if x != nil {
		return x.PlacementPartitionCount
	}
	return 0
}

type CreatePlacementGroupTaskState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePlacementGroupTaskState) Reset() {
	*x = CreatePlacementGroupTaskState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlacementGroupTaskState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlacementGroupTaskState) ProtoMessage() {}

func (x *CreatePlacementGroupTaskState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlacementGroupTaskState.ProtoReflect.Descriptor instead.
func (*CreatePlacementGroupTaskState) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescGZIP(), []int{1}
}

var File_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDesc = []byte{
	0x0a, 0x60, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x11, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x38, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x17, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22,
	0x1f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescData = file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDesc
)

func file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescData)
	})
	return file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDescData
}

var file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_goTypes = []interface{}{
	(*CreatePlacementGroupRequest)(nil),   // 0: placementgroup.CreatePlacementGroupRequest
	(*CreatePlacementGroupTaskState)(nil), // 1: placementgroup.CreatePlacementGroupTaskState
	(types.PlacementStrategy)(0),          // 2: types.PlacementStrategy
}
var file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_depIdxs = []int32{
	2, // 0: placementgroup.CreatePlacementGroupRequest.PlacementStrategy:type_name -> types.PlacementStrategy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_init()
}
func file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_init() {
	if File_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlacementGroupRequest); i {
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
		file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlacementGroupTaskState); i {
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
			RawDescriptor: file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_depIdxs,
		MessageInfos:      file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto = out.File
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_rawDesc = nil
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_goTypes = nil
	file_cloud_disk_manager_internal_pkg_services_placementgroup_protos_create_placement_group_task_proto_depIdxs = nil
}
