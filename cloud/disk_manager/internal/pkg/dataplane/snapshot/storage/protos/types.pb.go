// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/internal/pkg/dataplane/snapshot/storage/protos/types.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeletingSnapshotKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletingAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=DeletingAt,proto3" json:"DeletingAt,omitempty"`
	SnapshotId string                 `protobuf:"bytes,2,opt,name=SnapshotId,proto3" json:"SnapshotId,omitempty"`
}

func (x *DeletingSnapshotKey) Reset() {
	*x = DeletingSnapshotKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletingSnapshotKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletingSnapshotKey) ProtoMessage() {}

func (x *DeletingSnapshotKey) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletingSnapshotKey.ProtoReflect.Descriptor instead.
func (*DeletingSnapshotKey) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescGZIP(), []int{0}
}

func (x *DeletingSnapshotKey) GetDeletingAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletingAt
	}
	return nil
}

func (x *DeletingSnapshotKey) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

var File_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x3a,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescData = file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDesc
)

func file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescData)
	})
	return file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDescData
}

var file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_goTypes = []interface{}{
	(*DeletingSnapshotKey)(nil),   // 0: types.DeletingSnapshotKey
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_depIdxs = []int32{
	1, // 0: types.DeletingSnapshotKey.DeletingAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_init()
}
func file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_init() {
	if File_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletingSnapshotKey); i {
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
			RawDescriptor: file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_depIdxs,
		MessageInfos:      file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto = out.File
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_rawDesc = nil
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_goTypes = nil
	file_cloud_disk_manager_internal_pkg_dataplane_snapshot_storage_protos_types_proto_depIdxs = nil
}
