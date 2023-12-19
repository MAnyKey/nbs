// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/internal/pkg/services/images/protos/create_image_from_disk_task.proto

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

type CreateImageFromDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcDisk           *types.Disk       `protobuf:"bytes,1,opt,name=SrcDisk,proto3" json:"SrcDisk,omitempty"`
	DstImageId        string            `protobuf:"bytes,2,opt,name=DstImageId,proto3" json:"DstImageId,omitempty"`
	FolderId          string            `protobuf:"bytes,3,opt,name=FolderId,proto3" json:"FolderId,omitempty"`
	OperationCloudId  string            `protobuf:"bytes,4,opt,name=OperationCloudId,proto3" json:"OperationCloudId,omitempty"`
	OperationFolderId string            `protobuf:"bytes,5,opt,name=OperationFolderId,proto3" json:"OperationFolderId,omitempty"`
	DiskPools         []*types.DiskPool `protobuf:"bytes,6,rep,name=DiskPools,proto3" json:"DiskPools,omitempty"`
	UseDataplaneTasks bool              `protobuf:"varint,7,opt,name=UseDataplaneTasks,proto3" json:"UseDataplaneTasks,omitempty"`
	UseS3             bool              `protobuf:"varint,8,opt,name=UseS3,proto3" json:"UseS3,omitempty"`
}

func (x *CreateImageFromDiskRequest) Reset() {
	*x = CreateImageFromDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImageFromDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImageFromDiskRequest) ProtoMessage() {}

func (x *CreateImageFromDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImageFromDiskRequest.ProtoReflect.Descriptor instead.
func (*CreateImageFromDiskRequest) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateImageFromDiskRequest) GetSrcDisk() *types.Disk {
	if x != nil {
		return x.SrcDisk
	}
	return nil
}

func (x *CreateImageFromDiskRequest) GetDstImageId() string {
	if x != nil {
		return x.DstImageId
	}
	return ""
}

func (x *CreateImageFromDiskRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateImageFromDiskRequest) GetOperationCloudId() string {
	if x != nil {
		return x.OperationCloudId
	}
	return ""
}

func (x *CreateImageFromDiskRequest) GetOperationFolderId() string {
	if x != nil {
		return x.OperationFolderId
	}
	return ""
}

func (x *CreateImageFromDiskRequest) GetDiskPools() []*types.DiskPool {
	if x != nil {
		return x.DiskPools
	}
	return nil
}

func (x *CreateImageFromDiskRequest) GetUseDataplaneTasks() bool {
	if x != nil {
		return x.UseDataplaneTasks
	}
	return false
}

func (x *CreateImageFromDiskRequest) GetUseS3() bool {
	if x != nil {
		return x.UseS3
	}
	return false
}

type CreateImageFromDiskTaskState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset           int64   `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Progress         float64 `protobuf:"fixed64,3,opt,name=Progress,proto3" json:"Progress,omitempty"`
	ImageSize        int64   `protobuf:"varint,4,opt,name=ImageSize,proto3" json:"ImageSize,omitempty"`
	ImageStorageSize int64   `protobuf:"varint,5,opt,name=ImageStorageSize,proto3" json:"ImageStorageSize,omitempty"`
	DataplaneTaskID  string  `protobuf:"bytes,6,opt,name=DataplaneTaskID,proto3" json:"DataplaneTaskID,omitempty"`
}

func (x *CreateImageFromDiskTaskState) Reset() {
	*x = CreateImageFromDiskTaskState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImageFromDiskTaskState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImageFromDiskTaskState) ProtoMessage() {}

func (x *CreateImageFromDiskTaskState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImageFromDiskTaskState.ProtoReflect.Descriptor instead.
func (*CreateImageFromDiskTaskState) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateImageFromDiskTaskState) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CreateImageFromDiskTaskState) GetProgress() float64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *CreateImageFromDiskTaskState) GetImageSize() int64 {
	if x != nil {
		return x.ImageSize
	}
	return 0
}

func (x *CreateImageFromDiskTaskState) GetImageStorageSize() int64 {
	if x != nil {
		return x.ImageStorageSize
	}
	return 0
}

func (x *CreateImageFromDiskTaskState) GetDataplaneTaskID() string {
	if x != nil {
		return x.DataplaneTaskID
	}
	return ""
}

var File_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDesc = []byte{
	0x0a, 0x58, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x44, 0x69, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x73, 0x6b, 0x52, 0x07, 0x53, 0x72, 0x63, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x44,
	0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x44, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73,
	0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x55, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x53, 0x33, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x55,
	0x73, 0x65, 0x53, 0x33, 0x22, 0xc6, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x42, 0x54, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescData = file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDesc
)

func file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescData)
	})
	return file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDescData
}

var file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_goTypes = []interface{}{
	(*CreateImageFromDiskRequest)(nil),   // 0: images.CreateImageFromDiskRequest
	(*CreateImageFromDiskTaskState)(nil), // 1: images.CreateImageFromDiskTaskState
	(*types.Disk)(nil),                   // 2: types.Disk
	(*types.DiskPool)(nil),               // 3: types.DiskPool
}
var file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_depIdxs = []int32{
	2, // 0: images.CreateImageFromDiskRequest.SrcDisk:type_name -> types.Disk
	3, // 1: images.CreateImageFromDiskRequest.DiskPools:type_name -> types.DiskPool
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_init()
}
func file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_init() {
	if File_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImageFromDiskRequest); i {
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
		file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImageFromDiskTaskState); i {
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
			RawDescriptor: file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_depIdxs,
		MessageInfos:      file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto = out.File
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_rawDesc = nil
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_goTypes = nil
	file_cloud_disk_manager_internal_pkg_services_images_protos_create_image_from_disk_task_proto_depIdxs = nil
}
