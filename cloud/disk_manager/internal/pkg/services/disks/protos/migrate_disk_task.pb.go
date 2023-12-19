// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/disk_manager/internal/pkg/services/disks/protos/migrate_disk_task.proto

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

type MigrateDiskTaskEvents int32

const (
	MigrateDiskTaskEvents_FINISH_MIGRATION MigrateDiskTaskEvents = 0
)

// Enum value maps for MigrateDiskTaskEvents.
var (
	MigrateDiskTaskEvents_name = map[int32]string{
		0: "FINISH_MIGRATION",
	}
	MigrateDiskTaskEvents_value = map[string]int32{
		"FINISH_MIGRATION": 0,
	}
)

func (x MigrateDiskTaskEvents) Enum() *MigrateDiskTaskEvents {
	p := new(MigrateDiskTaskEvents)
	*p = x
	return p
}

func (x MigrateDiskTaskEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrateDiskTaskEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes[0].Descriptor()
}

func (MigrateDiskTaskEvents) Type() protoreflect.EnumType {
	return &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes[0]
}

func (x MigrateDiskTaskEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrateDiskTaskEvents.Descriptor instead.
func (MigrateDiskTaskEvents) EnumDescriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP(), []int{0}
}

// Should be in sync with statuses from public API.
type MigrationStatus int32

const (
	MigrationStatus_Unspecified          MigrationStatus = 0
	MigrationStatus_Replicating          MigrationStatus = 1
	MigrationStatus_FinishingReplication MigrationStatus = 2
	MigrationStatus_ReplicationFinished  MigrationStatus = 3
	MigrationStatus_Finishing            MigrationStatus = 4
)

// Enum value maps for MigrationStatus.
var (
	MigrationStatus_name = map[int32]string{
		0: "Unspecified",
		1: "Replicating",
		2: "FinishingReplication",
		3: "ReplicationFinished",
		4: "Finishing",
	}
	MigrationStatus_value = map[string]int32{
		"Unspecified":          0,
		"Replicating":          1,
		"FinishingReplication": 2,
		"ReplicationFinished":  3,
		"Finishing":            4,
	}
)

func (x MigrationStatus) Enum() *MigrationStatus {
	p := new(MigrationStatus)
	*p = x
	return p
}

func (x MigrationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes[1].Descriptor()
}

func (MigrationStatus) Type() protoreflect.EnumType {
	return &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes[1]
}

func (x MigrationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationStatus.Descriptor instead.
func (MigrationStatus) EnumDescriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP(), []int{1}
}

type MigrateDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disk                       *types.Disk `protobuf:"bytes,1,opt,name=Disk,proto3" json:"Disk,omitempty"`
	DstZoneId                  string      `protobuf:"bytes,2,opt,name=DstZoneId,proto3" json:"DstZoneId,omitempty"`
	DstPlacementGroupId        string      `protobuf:"bytes,3,opt,name=DstPlacementGroupId,proto3" json:"DstPlacementGroupId,omitempty"`
	DstPlacementPartitionIndex uint32      `protobuf:"varint,4,opt,name=DstPlacementPartitionIndex,proto3" json:"DstPlacementPartitionIndex,omitempty"`
}

func (x *MigrateDiskRequest) Reset() {
	*x = MigrateDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateDiskRequest) ProtoMessage() {}

func (x *MigrateDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateDiskRequest.ProtoReflect.Descriptor instead.
func (*MigrateDiskRequest) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP(), []int{0}
}

func (x *MigrateDiskRequest) GetDisk() *types.Disk {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *MigrateDiskRequest) GetDstZoneId() string {
	if x != nil {
		return x.DstZoneId
	}
	return ""
}

func (x *MigrateDiskRequest) GetDstPlacementGroupId() string {
	if x != nil {
		return x.DstPlacementGroupId
	}
	return ""
}

func (x *MigrateDiskRequest) GetDstPlacementPartitionIndex() uint32 {
	if x != nil {
		return x.DstPlacementPartitionIndex
	}
	return 0
}

type RelocateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseDiskID       string `protobuf:"bytes,1,opt,name=BaseDiskID,proto3" json:"BaseDiskID,omitempty"`
	TargetBaseDiskID string `protobuf:"bytes,2,opt,name=TargetBaseDiskID,proto3" json:"TargetBaseDiskID,omitempty"`
	SlotGeneration   uint64 `protobuf:"varint,3,opt,name=SlotGeneration,proto3" json:"SlotGeneration,omitempty"`
}

func (x *RelocateInfo) Reset() {
	*x = RelocateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelocateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelocateInfo) ProtoMessage() {}

func (x *RelocateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelocateInfo.ProtoReflect.Descriptor instead.
func (*RelocateInfo) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP(), []int{1}
}

func (x *RelocateInfo) GetBaseDiskID() string {
	if x != nil {
		return x.BaseDiskID
	}
	return ""
}

func (x *RelocateInfo) GetTargetBaseDiskID() string {
	if x != nil {
		return x.TargetBaseDiskID
	}
	return ""
}

func (x *RelocateInfo) GetSlotGeneration() uint64 {
	if x != nil {
		return x.SlotGeneration
	}
	return 0
}

type MigrateDiskTaskState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDiskCloned    bool            `protobuf:"varint,6,opt,name=IsDiskCloned,proto3" json:"IsDiskCloned,omitempty"`
	FillSeqNumber   uint64          `protobuf:"varint,7,opt,name=FillSeqNumber,proto3" json:"FillSeqNumber,omitempty"`
	FillGeneration  uint64          `protobuf:"varint,8,opt,name=FillGeneration,proto3" json:"FillGeneration,omitempty"`
	ReplicateTaskID string          `protobuf:"bytes,9,opt,name=ReplicateTaskID,proto3" json:"ReplicateTaskID,omitempty"`
	Status          MigrationStatus `protobuf:"varint,10,opt,name=Status,proto3,enum=disks.MigrationStatus" json:"Status,omitempty"`
	RelocateInfo    *RelocateInfo   `protobuf:"bytes,11,opt,name=RelocateInfo,proto3" json:"RelocateInfo,omitempty"`
}

func (x *MigrateDiskTaskState) Reset() {
	*x = MigrateDiskTaskState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateDiskTaskState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateDiskTaskState) ProtoMessage() {}

func (x *MigrateDiskTaskState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateDiskTaskState.ProtoReflect.Descriptor instead.
func (*MigrateDiskTaskState) Descriptor() ([]byte, []int) {
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP(), []int{2}
}

func (x *MigrateDiskTaskState) GetIsDiskCloned() bool {
	if x != nil {
		return x.IsDiskCloned
	}
	return false
}

func (x *MigrateDiskTaskState) GetFillSeqNumber() uint64 {
	if x != nil {
		return x.FillSeqNumber
	}
	return 0
}

func (x *MigrateDiskTaskState) GetFillGeneration() uint64 {
	if x != nil {
		return x.FillGeneration
	}
	return 0
}

func (x *MigrateDiskTaskState) GetReplicateTaskID() string {
	if x != nil {
		return x.ReplicateTaskID
	}
	return ""
}

func (x *MigrateDiskTaskState) GetStatus() MigrationStatus {
	if x != nil {
		return x.Status
	}
	return MigrationStatus_Unspecified
}

func (x *MigrateDiskTaskState) GetRelocateInfo() *RelocateInfo {
	if x != nil {
		return x.RelocateInfo
	}
	return nil
}

var File_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto protoreflect.FileDescriptor

var file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x12, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x04, 0x44, 0x69, 0x73,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x13, 0x44, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x44, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x1a, 0x44, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x44, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x6b,
	0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x44, 0x69, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x26,
	0x0a, 0x0e, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x02, 0x0a, 0x14, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x49, 0x73, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49, 0x73, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x6c, 0x6f,
	0x6e, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x71, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x46, 0x69, 0x6c, 0x6c,
	0x53, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x6c,
	0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x46, 0x69, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x52,
	0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x52, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x06, 0x2a, 0x2d, 0x0a, 0x15, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x4d, 0x49,
	0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x2a, 0x75, 0x0a, 0x0f, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescOnce sync.Once
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescData = file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDesc
)

func file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescGZIP() []byte {
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescOnce.Do(func() {
		file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescData)
	})
	return file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDescData
}

var file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_goTypes = []interface{}{
	(MigrateDiskTaskEvents)(0),   // 0: disks.MigrateDiskTaskEvents
	(MigrationStatus)(0),         // 1: disks.MigrationStatus
	(*MigrateDiskRequest)(nil),   // 2: disks.MigrateDiskRequest
	(*RelocateInfo)(nil),         // 3: disks.RelocateInfo
	(*MigrateDiskTaskState)(nil), // 4: disks.MigrateDiskTaskState
	(*types.Disk)(nil),           // 5: types.Disk
}
var file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_depIdxs = []int32{
	5, // 0: disks.MigrateDiskRequest.Disk:type_name -> types.Disk
	1, // 1: disks.MigrateDiskTaskState.Status:type_name -> disks.MigrationStatus
	3, // 2: disks.MigrateDiskTaskState.RelocateInfo:type_name -> disks.RelocateInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_init()
}
func file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_init() {
	if File_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateDiskRequest); i {
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
		file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelocateInfo); i {
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
		file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateDiskTaskState); i {
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
			RawDescriptor: file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_goTypes,
		DependencyIndexes: file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_depIdxs,
		EnumInfos:         file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_enumTypes,
		MessageInfos:      file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_msgTypes,
	}.Build()
	File_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto = out.File
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_rawDesc = nil
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_goTypes = nil
	file_cloud_disk_manager_internal_pkg_services_disks_protos_migrate_disk_task_proto_depIdxs = nil
}
