// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/blockstore/public/api/protos/client.proto

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

type EHostType int32

const (
	EHostType_HOST_TYPE_DEFAULT   EHostType = 0
	EHostType_HOST_TYPE_DEDICATED EHostType = 1
)

// Enum value maps for EHostType.
var (
	EHostType_name = map[int32]string{
		0: "HOST_TYPE_DEFAULT",
		1: "HOST_TYPE_DEDICATED",
	}
	EHostType_value = map[string]int32{
		"HOST_TYPE_DEFAULT":   0,
		"HOST_TYPE_DEDICATED": 1,
	}
)

func (x EHostType) Enum() *EHostType {
	p := new(EHostType)
	*p = x
	return p
}

func (x EHostType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EHostType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_blockstore_public_api_protos_client_proto_enumTypes[0].Descriptor()
}

func (EHostType) Type() protoreflect.EnumType {
	return &file_cloud_blockstore_public_api_protos_client_proto_enumTypes[0]
}

func (x EHostType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EHostType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EHostType(num)
	return nil
}

// Deprecated: Use EHostType.Descriptor instead.
func (EHostType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_client_proto_rawDescGZIP(), []int{0}
}

type TClientMediaKindPerformanceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read bytes per sec.
	MaxReadBandwidth *uint32 `protobuf:"varint,1,opt,name=MaxReadBandwidth" json:"MaxReadBandwidth,omitempty"`
	// Read iops.
	MaxReadIops *uint32 `protobuf:"varint,2,opt,name=MaxReadIops" json:"MaxReadIops,omitempty"`
	// Write bytes per sec.
	MaxWriteBandwidth *uint32 `protobuf:"varint,3,opt,name=MaxWriteBandwidth" json:"MaxWriteBandwidth,omitempty"`
	// Write iops.
	MaxWriteIops *uint32 `protobuf:"varint,4,opt,name=MaxWriteIops" json:"MaxWriteIops,omitempty"`
}

func (x *TClientMediaKindPerformanceProfile) Reset() {
	*x = TClientMediaKindPerformanceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TClientMediaKindPerformanceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientMediaKindPerformanceProfile) ProtoMessage() {}

func (x *TClientMediaKindPerformanceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientMediaKindPerformanceProfile.ProtoReflect.Descriptor instead.
func (*TClientMediaKindPerformanceProfile) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_client_proto_rawDescGZIP(), []int{0}
}

func (x *TClientMediaKindPerformanceProfile) GetMaxReadBandwidth() uint32 {
	if x != nil && x.MaxReadBandwidth != nil {
		return *x.MaxReadBandwidth
	}
	return 0
}

func (x *TClientMediaKindPerformanceProfile) GetMaxReadIops() uint32 {
	if x != nil && x.MaxReadIops != nil {
		return *x.MaxReadIops
	}
	return 0
}

func (x *TClientMediaKindPerformanceProfile) GetMaxWriteBandwidth() uint32 {
	if x != nil && x.MaxWriteBandwidth != nil {
		return *x.MaxWriteBandwidth
	}
	return 0
}

func (x *TClientMediaKindPerformanceProfile) GetMaxWriteIops() uint32 {
	if x != nil && x.MaxWriteIops != nil {
		return *x.MaxWriteIops
	}
	return 0
}

type TClientPerformanceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Performance profile used to throttle requests to network-hdd volumes.
	HDDProfile *TClientMediaKindPerformanceProfile `protobuf:"bytes,1,opt,name=HDDProfile" json:"HDDProfile,omitempty"`
	// Performance profile used to throttle requests to network-ssd volumes.
	SSDProfile *TClientMediaKindPerformanceProfile `protobuf:"bytes,2,opt,name=SSDProfile" json:"SSDProfile,omitempty"`
	// Performance profile used to throttle requests to
	// network-ssd-nonreplicated volumes.
	NonreplProfile *TClientMediaKindPerformanceProfile `protobuf:"bytes,3,opt,name=NonreplProfile" json:"NonreplProfile,omitempty"`
	// Performance profile used to throttle requests to network-ssd-mirror2
	// volumes.
	Mirror2Profile *TClientMediaKindPerformanceProfile `protobuf:"bytes,4,opt,name=Mirror2Profile" json:"Mirror2Profile,omitempty"`
	// Performance profile used to throttle requests to network-ssd-mirror3
	// volumes.
	Mirror3Profile *TClientMediaKindPerformanceProfile `protobuf:"bytes,5,opt,name=Mirror3Profile" json:"Mirror3Profile,omitempty"`
	// Maximum burst length in milliseconds.
	BurstTime *uint32 `protobuf:"varint,6,opt,name=BurstTime" json:"BurstTime,omitempty"`
	// Performance profile used to throttle requests to
	// network-hdd-nonreplicated volumes.
	HddNonreplProfile *TClientMediaKindPerformanceProfile `protobuf:"bytes,7,opt,name=HddNonreplProfile" json:"HddNonreplProfile,omitempty"`
}

func (x *TClientPerformanceProfile) Reset() {
	*x = TClientPerformanceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TClientPerformanceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientPerformanceProfile) ProtoMessage() {}

func (x *TClientPerformanceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientPerformanceProfile.ProtoReflect.Descriptor instead.
func (*TClientPerformanceProfile) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_client_proto_rawDescGZIP(), []int{1}
}

func (x *TClientPerformanceProfile) GetHDDProfile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.HDDProfile
	}
	return nil
}

func (x *TClientPerformanceProfile) GetSSDProfile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.SSDProfile
	}
	return nil
}

func (x *TClientPerformanceProfile) GetNonreplProfile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.NonreplProfile
	}
	return nil
}

func (x *TClientPerformanceProfile) GetMirror2Profile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.Mirror2Profile
	}
	return nil
}

func (x *TClientPerformanceProfile) GetMirror3Profile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.Mirror3Profile
	}
	return nil
}

func (x *TClientPerformanceProfile) GetBurstTime() uint32 {
	if x != nil && x.BurstTime != nil {
		return *x.BurstTime
	}
	return 0
}

func (x *TClientPerformanceProfile) GetHddNonreplProfile() *TClientMediaKindPerformanceProfile {
	if x != nil {
		return x.HddNonreplProfile
	}
	return nil
}

type TClientProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. the number of cores allocated to the VM.
	CpuUnitCount *uint32 `protobuf:"varint,1,opt,name=CpuUnitCount" json:"CpuUnitCount,omitempty"`
	// Throttling policy depends on the type of the host.
	HostType *EHostType `protobuf:"varint,2,opt,name=HostType,enum=NCloud.NBlockStore.NProto.EHostType" json:"HostType,omitempty"`
}

func (x *TClientProfile) Reset() {
	*x = TClientProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TClientProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientProfile) ProtoMessage() {}

func (x *TClientProfile) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_public_api_protos_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientProfile.ProtoReflect.Descriptor instead.
func (*TClientProfile) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_client_proto_rawDescGZIP(), []int{2}
}

func (x *TClientProfile) GetCpuUnitCount() uint32 {
	if x != nil && x.CpuUnitCount != nil {
		return *x.CpuUnitCount
	}
	return 0
}

func (x *TClientProfile) GetHostType() EHostType {
	if x != nil && x.HostType != nil {
		return *x.HostType
	}
	return EHostType_HOST_TYPE_DEFAULT
}

var File_cloud_blockstore_public_api_protos_client_proto protoreflect.FileDescriptor

var file_cloud_blockstore_public_api_protos_client_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a,
	0x22, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69, 0x6e,
	0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6f, 0x70, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6f, 0x70,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x61, 0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x4d, 0x61,
	0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x6f, 0x70, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x4d, 0x61, 0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49,
	0x6f, 0x70, 0x73, 0x22, 0x99, 0x05, 0x0a, 0x19, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x48, 0x44, 0x44, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69,
	0x6e, 0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x0a, 0x48, 0x44, 0x44, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x5d, 0x0a, 0x0a, 0x53, 0x53, 0x44, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69, 0x6e,
	0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x0a, 0x53, 0x53, 0x44, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x65, 0x0a, 0x0e, 0x4e, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x4b, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0e, 0x4e, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6c, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0e, 0x4d,
	0x69, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x65, 0x0a,
	0x0e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x33, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69,
	0x6e, 0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x0e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x33, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75, 0x72, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42, 0x75, 0x72, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x11, 0x48, 0x64, 0x64, 0x4e, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6c,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x11, 0x48, 0x64,
	0x64, 0x4e, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x76, 0x0a, 0x0e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x43, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x48,
	0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x3b, 0x0a, 0x09, 0x45, 0x48, 0x6f, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x48,
	0x4f, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x44, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_cloud_blockstore_public_api_protos_client_proto_rawDescOnce sync.Once
	file_cloud_blockstore_public_api_protos_client_proto_rawDescData = file_cloud_blockstore_public_api_protos_client_proto_rawDesc
)

func file_cloud_blockstore_public_api_protos_client_proto_rawDescGZIP() []byte {
	file_cloud_blockstore_public_api_protos_client_proto_rawDescOnce.Do(func() {
		file_cloud_blockstore_public_api_protos_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_blockstore_public_api_protos_client_proto_rawDescData)
	})
	return file_cloud_blockstore_public_api_protos_client_proto_rawDescData
}

var file_cloud_blockstore_public_api_protos_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_blockstore_public_api_protos_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_blockstore_public_api_protos_client_proto_goTypes = []interface{}{
	(EHostType)(0), // 0: NCloud.NBlockStore.NProto.EHostType
	(*TClientMediaKindPerformanceProfile)(nil), // 1: NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	(*TClientPerformanceProfile)(nil),          // 2: NCloud.NBlockStore.NProto.TClientPerformanceProfile
	(*TClientProfile)(nil),                     // 3: NCloud.NBlockStore.NProto.TClientProfile
}
var file_cloud_blockstore_public_api_protos_client_proto_depIdxs = []int32{
	1, // 0: NCloud.NBlockStore.NProto.TClientPerformanceProfile.HDDProfile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	1, // 1: NCloud.NBlockStore.NProto.TClientPerformanceProfile.SSDProfile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	1, // 2: NCloud.NBlockStore.NProto.TClientPerformanceProfile.NonreplProfile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	1, // 3: NCloud.NBlockStore.NProto.TClientPerformanceProfile.Mirror2Profile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	1, // 4: NCloud.NBlockStore.NProto.TClientPerformanceProfile.Mirror3Profile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	1, // 5: NCloud.NBlockStore.NProto.TClientPerformanceProfile.HddNonreplProfile:type_name -> NCloud.NBlockStore.NProto.TClientMediaKindPerformanceProfile
	0, // 6: NCloud.NBlockStore.NProto.TClientProfile.HostType:type_name -> NCloud.NBlockStore.NProto.EHostType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_blockstore_public_api_protos_client_proto_init() }
func file_cloud_blockstore_public_api_protos_client_proto_init() {
	if File_cloud_blockstore_public_api_protos_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_blockstore_public_api_protos_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TClientMediaKindPerformanceProfile); i {
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
		file_cloud_blockstore_public_api_protos_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TClientPerformanceProfile); i {
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
		file_cloud_blockstore_public_api_protos_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TClientProfile); i {
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
			RawDescriptor: file_cloud_blockstore_public_api_protos_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_blockstore_public_api_protos_client_proto_goTypes,
		DependencyIndexes: file_cloud_blockstore_public_api_protos_client_proto_depIdxs,
		EnumInfos:         file_cloud_blockstore_public_api_protos_client_proto_enumTypes,
		MessageInfos:      file_cloud_blockstore_public_api_protos_client_proto_msgTypes,
	}.Build()
	File_cloud_blockstore_public_api_protos_client_proto = out.File
	file_cloud_blockstore_public_api_protos_client_proto_rawDesc = nil
	file_cloud_blockstore_public_api_protos_client_proto_goTypes = nil
	file_cloud_blockstore_public_api_protos_client_proto_depIdxs = nil
}
