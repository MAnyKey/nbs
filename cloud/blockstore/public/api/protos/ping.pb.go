// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/blockstore/public/api/protos/ping.proto

package protos

import (
	protos "github.com/ydb-platform/nbs/cloud/storage/core/protos"
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

type TPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional request headers.
	Headers *THeaders `protobuf:"bytes,1,opt,name=Headers,proto3" json:"Headers,omitempty"`
}

func (x *TPingRequest) Reset() {
	*x = TPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPingRequest) ProtoMessage() {}

func (x *TPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPingRequest.ProtoReflect.Descriptor instead.
func (*TPingRequest) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_ping_proto_rawDescGZIP(), []int{0}
}

func (x *TPingRequest) GetHeaders() *THeaders {
	if x != nil {
		return x.Headers
	}
	return nil
}

type TPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional error, set only if error happened.
	Error *protos.TError `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	// Most recent stats, will be used for load balancing
	LastByteCount    uint64 `protobuf:"varint,2,opt,name=LastByteCount,proto3" json:"LastByteCount,omitempty"`
	LastRequestCount uint64 `protobuf:"varint,3,opt,name=LastRequestCount,proto3" json:"LastRequestCount,omitempty"`
}

func (x *TPingResponse) Reset() {
	*x = TPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPingResponse) ProtoMessage() {}

func (x *TPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPingResponse.ProtoReflect.Descriptor instead.
func (*TPingResponse) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_public_api_protos_ping_proto_rawDescGZIP(), []int{1}
}

func (x *TPingResponse) GetError() *protos.TError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TPingResponse) GetLastByteCount() uint64 {
	if x != nil {
		return x.LastByteCount
	}
	return 0
}

func (x *TPingResponse) GetLastRequestCount() uint64 {
	if x != nil {
		return x.LastRequestCount
	}
	return 0
}

var File_cloud_blockstore_public_api_protos_ping_proto protoreflect.FileDescriptor

var file_cloud_blockstore_public_api_protos_ping_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0c, 0x54, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x54, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e,
	0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_blockstore_public_api_protos_ping_proto_rawDescOnce sync.Once
	file_cloud_blockstore_public_api_protos_ping_proto_rawDescData = file_cloud_blockstore_public_api_protos_ping_proto_rawDesc
)

func file_cloud_blockstore_public_api_protos_ping_proto_rawDescGZIP() []byte {
	file_cloud_blockstore_public_api_protos_ping_proto_rawDescOnce.Do(func() {
		file_cloud_blockstore_public_api_protos_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_blockstore_public_api_protos_ping_proto_rawDescData)
	})
	return file_cloud_blockstore_public_api_protos_ping_proto_rawDescData
}

var file_cloud_blockstore_public_api_protos_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_blockstore_public_api_protos_ping_proto_goTypes = []interface{}{
	(*TPingRequest)(nil),  // 0: NCloud.NBlockStore.NProto.TPingRequest
	(*TPingResponse)(nil), // 1: NCloud.NBlockStore.NProto.TPingResponse
	(*THeaders)(nil),      // 2: NCloud.NBlockStore.NProto.THeaders
	(*protos.TError)(nil), // 3: NCloud.NProto.TError
}
var file_cloud_blockstore_public_api_protos_ping_proto_depIdxs = []int32{
	2, // 0: NCloud.NBlockStore.NProto.TPingRequest.Headers:type_name -> NCloud.NBlockStore.NProto.THeaders
	3, // 1: NCloud.NBlockStore.NProto.TPingResponse.Error:type_name -> NCloud.NProto.TError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_blockstore_public_api_protos_ping_proto_init() }
func file_cloud_blockstore_public_api_protos_ping_proto_init() {
	if File_cloud_blockstore_public_api_protos_ping_proto != nil {
		return
	}
	file_cloud_blockstore_public_api_protos_headers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TPingRequest); i {
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
		file_cloud_blockstore_public_api_protos_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TPingResponse); i {
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
			RawDescriptor: file_cloud_blockstore_public_api_protos_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_blockstore_public_api_protos_ping_proto_goTypes,
		DependencyIndexes: file_cloud_blockstore_public_api_protos_ping_proto_depIdxs,
		MessageInfos:      file_cloud_blockstore_public_api_protos_ping_proto_msgTypes,
	}.Build()
	File_cloud_blockstore_public_api_protos_ping_proto = out.File
	file_cloud_blockstore_public_api_protos_ping_proto_rawDesc = nil
	file_cloud_blockstore_public_api_protos_ping_proto_goTypes = nil
	file_cloud_blockstore_public_api_protos_ping_proto_depIdxs = nil
}
