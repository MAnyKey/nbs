// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/filestore/config/server.proto

package config

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

type TServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host name or address to listen on.
	Host *string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	// Port to listen on.
	Port *uint32 `protobuf:"varint,2,opt,name=Port" json:"Port,omitempty"`
	// Number of gRPC threads.
	ThreadsCount *uint32 `protobuf:"varint,3,opt,name=ThreadsCount" json:"ThreadsCount,omitempty"`
	// Limit of grpc threads count.
	GrpcThreadsLimit *uint32 `protobuf:"varint,4,opt,name=GrpcThreadsLimit" json:"GrpcThreadsLimit,omitempty"`
	// Maximum size of gRPC message.
	MaxMessageSize *uint32 `protobuf:"varint,5,opt,name=MaxMessageSize" json:"MaxMessageSize,omitempty"`
	// Memory quota for gRPC request processing.
	MemoryQuotaBytes *uint32 `protobuf:"varint,6,opt,name=MemoryQuotaBytes" json:"MemoryQuotaBytes,omitempty"`
	// Number of prepared gRPC requests.
	PreparedRequestsCount *uint32 `protobuf:"varint,7,opt,name=PreparedRequestsCount" json:"PreparedRequestsCount,omitempty"`
	// KeepAlive details.
	KeepAliveEnabled      *bool   `protobuf:"varint,8,opt,name=KeepAliveEnabled" json:"KeepAliveEnabled,omitempty"`
	KeepAliveIdleTimeout  *uint32 `protobuf:"varint,9,opt,name=KeepAliveIdleTimeout" json:"KeepAliveIdleTimeout,omitempty"`    // (in milliseconds).
	KeepAliveProbeTimeout *uint32 `protobuf:"varint,10,opt,name=KeepAliveProbeTimeout" json:"KeepAliveProbeTimeout,omitempty"` // (in milliseconds).
	KeepAliveProbesCount  *uint32 `protobuf:"varint,11,opt,name=KeepAliveProbesCount" json:"KeepAliveProbesCount,omitempty"`
	ShutdownTimeout       *uint32 `protobuf:"varint,12,opt,name=ShutdownTimeout" json:"ShutdownTimeout,omitempty"` // (in milliseconds).
	// Host name or address to listen on (with TLS enabled).
	SecureHost *string `protobuf:"bytes,13,opt,name=SecureHost" json:"SecureHost,omitempty"`
	// Port to listen on (with TLS enabled).
	SecurePort *uint32 `protobuf:"varint,14,opt,name=SecurePort" json:"SecurePort,omitempty"`
	// TLS details.
	RootCertsFile *string                `protobuf:"bytes,15,opt,name=RootCertsFile" json:"RootCertsFile,omitempty"`
	Certs         []*protos.TCertificate `protobuf:"bytes,16,rep,name=Certs" json:"Certs,omitempty"`
	// Unix-socket details.
	UnixSocketPath    *string `protobuf:"bytes,17,opt,name=UnixSocketPath" json:"UnixSocketPath,omitempty"`
	UnixSocketBacklog *uint32 `protobuf:"varint,18,opt,name=UnixSocketBacklog" json:"UnixSocketBacklog,omitempty"`
}

func (x *TServerConfig) Reset() {
	*x = TServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TServerConfig) ProtoMessage() {}

func (x *TServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TServerConfig.ProtoReflect.Descriptor instead.
func (*TServerConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_server_proto_rawDescGZIP(), []int{0}
}

func (x *TServerConfig) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *TServerConfig) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *TServerConfig) GetThreadsCount() uint32 {
	if x != nil && x.ThreadsCount != nil {
		return *x.ThreadsCount
	}
	return 0
}

func (x *TServerConfig) GetGrpcThreadsLimit() uint32 {
	if x != nil && x.GrpcThreadsLimit != nil {
		return *x.GrpcThreadsLimit
	}
	return 0
}

func (x *TServerConfig) GetMaxMessageSize() uint32 {
	if x != nil && x.MaxMessageSize != nil {
		return *x.MaxMessageSize
	}
	return 0
}

func (x *TServerConfig) GetMemoryQuotaBytes() uint32 {
	if x != nil && x.MemoryQuotaBytes != nil {
		return *x.MemoryQuotaBytes
	}
	return 0
}

func (x *TServerConfig) GetPreparedRequestsCount() uint32 {
	if x != nil && x.PreparedRequestsCount != nil {
		return *x.PreparedRequestsCount
	}
	return 0
}

func (x *TServerConfig) GetKeepAliveEnabled() bool {
	if x != nil && x.KeepAliveEnabled != nil {
		return *x.KeepAliveEnabled
	}
	return false
}

func (x *TServerConfig) GetKeepAliveIdleTimeout() uint32 {
	if x != nil && x.KeepAliveIdleTimeout != nil {
		return *x.KeepAliveIdleTimeout
	}
	return 0
}

func (x *TServerConfig) GetKeepAliveProbeTimeout() uint32 {
	if x != nil && x.KeepAliveProbeTimeout != nil {
		return *x.KeepAliveProbeTimeout
	}
	return 0
}

func (x *TServerConfig) GetKeepAliveProbesCount() uint32 {
	if x != nil && x.KeepAliveProbesCount != nil {
		return *x.KeepAliveProbesCount
	}
	return 0
}

func (x *TServerConfig) GetShutdownTimeout() uint32 {
	if x != nil && x.ShutdownTimeout != nil {
		return *x.ShutdownTimeout
	}
	return 0
}

func (x *TServerConfig) GetSecureHost() string {
	if x != nil && x.SecureHost != nil {
		return *x.SecureHost
	}
	return ""
}

func (x *TServerConfig) GetSecurePort() uint32 {
	if x != nil && x.SecurePort != nil {
		return *x.SecurePort
	}
	return 0
}

func (x *TServerConfig) GetRootCertsFile() string {
	if x != nil && x.RootCertsFile != nil {
		return *x.RootCertsFile
	}
	return ""
}

func (x *TServerConfig) GetCerts() []*protos.TCertificate {
	if x != nil {
		return x.Certs
	}
	return nil
}

func (x *TServerConfig) GetUnixSocketPath() string {
	if x != nil && x.UnixSocketPath != nil {
		return *x.UnixSocketPath
	}
	return ""
}

func (x *TServerConfig) GetUnixSocketBacklog() uint32 {
	if x != nil && x.UnixSocketBacklog != nil {
		return *x.UnixSocketBacklog
	}
	return 0
}

type TNullServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TNullServiceConfig) Reset() {
	*x = TNullServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNullServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNullServiceConfig) ProtoMessage() {}

func (x *TNullServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNullServiceConfig.ProtoReflect.Descriptor instead.
func (*TNullServiceConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_server_proto_rawDescGZIP(), []int{1}
}

type TLocalServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootPath           *string `protobuf:"bytes,1,opt,name=RootPath" json:"RootPath,omitempty"`
	PathPrefix         *string `protobuf:"bytes,2,opt,name=PathPrefix" json:"PathPrefix,omitempty"`
	DefaultPermissions *uint32 `protobuf:"varint,3,opt,name=DefaultPermissions" json:"DefaultPermissions,omitempty"`
	IdleSessionTimeout *uint32 `protobuf:"varint,4,opt,name=IdleSessionTimeout" json:"IdleSessionTimeout,omitempty"`
	NumThreads         *uint32 `protobuf:"varint,5,opt,name=NumThreads" json:"NumThreads,omitempty"`
}

func (x *TLocalServiceConfig) Reset() {
	*x = TLocalServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLocalServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLocalServiceConfig) ProtoMessage() {}

func (x *TLocalServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLocalServiceConfig.ProtoReflect.Descriptor instead.
func (*TLocalServiceConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_server_proto_rawDescGZIP(), []int{2}
}

func (x *TLocalServiceConfig) GetRootPath() string {
	if x != nil && x.RootPath != nil {
		return *x.RootPath
	}
	return ""
}

func (x *TLocalServiceConfig) GetPathPrefix() string {
	if x != nil && x.PathPrefix != nil {
		return *x.PathPrefix
	}
	return ""
}

func (x *TLocalServiceConfig) GetDefaultPermissions() uint32 {
	if x != nil && x.DefaultPermissions != nil {
		return *x.DefaultPermissions
	}
	return 0
}

func (x *TLocalServiceConfig) GetIdleSessionTimeout() uint32 {
	if x != nil && x.IdleSessionTimeout != nil {
		return *x.IdleSessionTimeout
	}
	return 0
}

func (x *TLocalServiceConfig) GetNumThreads() uint32 {
	if x != nil && x.NumThreads != nil {
		return *x.NumThreads
	}
	return 0
}

type TKikimrServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TKikimrServiceConfig) Reset() {
	*x = TKikimrServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TKikimrServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TKikimrServiceConfig) ProtoMessage() {}

func (x *TKikimrServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TKikimrServiceConfig.ProtoReflect.Descriptor instead.
func (*TKikimrServiceConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_server_proto_rawDescGZIP(), []int{3}
}

type TServerAppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerConfig        *TServerConfig        `protobuf:"bytes,1,opt,name=ServerConfig" json:"ServerConfig,omitempty"`
	NullServiceConfig   *TNullServiceConfig   `protobuf:"bytes,2,opt,name=NullServiceConfig" json:"NullServiceConfig,omitempty"`
	LocalServiceConfig  *TLocalServiceConfig  `protobuf:"bytes,3,opt,name=LocalServiceConfig" json:"LocalServiceConfig,omitempty"`
	KikimrServiceConfig *TKikimrServiceConfig `protobuf:"bytes,4,opt,name=KikimrServiceConfig" json:"KikimrServiceConfig,omitempty"`
}

func (x *TServerAppConfig) Reset() {
	*x = TServerAppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TServerAppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TServerAppConfig) ProtoMessage() {}

func (x *TServerAppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TServerAppConfig.ProtoReflect.Descriptor instead.
func (*TServerAppConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_server_proto_rawDescGZIP(), []int{4}
}

func (x *TServerAppConfig) GetServerConfig() *TServerConfig {
	if x != nil {
		return x.ServerConfig
	}
	return nil
}

func (x *TServerAppConfig) GetNullServiceConfig() *TNullServiceConfig {
	if x != nil {
		return x.NullServiceConfig
	}
	return nil
}

func (x *TServerAppConfig) GetLocalServiceConfig() *TLocalServiceConfig {
	if x != nil {
		return x.LocalServiceConfig
	}
	return nil
}

func (x *TServerAppConfig) GetKikimrServiceConfig() *TKikimrServiceConfig {
	if x != nil {
		return x.KikimrServiceConfig
	}
	return nil
}

var File_cloud_filestore_config_server_proto protoreflect.FileDescriptor

var file_cloud_filestore_config_server_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x05, 0x0a,
	0x0d, 0x54, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x72,
	0x70, 0x63, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x47, 0x72, 0x70, 0x63, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x4b, 0x65, 0x65, 0x70,
	0x41, 0x6c, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x14,
	0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x4b, 0x65, 0x65, 0x70,
	0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x34, 0x0a, 0x15, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x6f, 0x6f,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x43, 0x65,
	0x72, 0x74, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x11, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b,
	0x6c, 0x6f, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xd1, 0x01, 0x0a, 0x13, 0x54, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2e, 0x0a,
	0x12, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a,
	0x12, 0x49, 0x64, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x49, 0x64, 0x6c, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x4e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x4e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x22, 0x16, 0x0a,
	0x14, 0x54, 0x4b, 0x69, 0x6b, 0x69, 0x6d, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xfc, 0x02, 0x0a, 0x10, 0x54, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x11, 0x4e, 0x75, 0x6c, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x4e,
	0x75, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x11, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x5d, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x60, 0x0a, 0x13, 0x4b, 0x69, 0x6b, 0x69, 0x6d, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x4b, 0x69, 0x6b, 0x69,
	0x6d, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x13, 0x4b, 0x69, 0x6b, 0x69, 0x6d, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
}

var (
	file_cloud_filestore_config_server_proto_rawDescOnce sync.Once
	file_cloud_filestore_config_server_proto_rawDescData = file_cloud_filestore_config_server_proto_rawDesc
)

func file_cloud_filestore_config_server_proto_rawDescGZIP() []byte {
	file_cloud_filestore_config_server_proto_rawDescOnce.Do(func() {
		file_cloud_filestore_config_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_filestore_config_server_proto_rawDescData)
	})
	return file_cloud_filestore_config_server_proto_rawDescData
}

var file_cloud_filestore_config_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_filestore_config_server_proto_goTypes = []interface{}{
	(*TServerConfig)(nil),        // 0: NCloud.NFileStore.NProto.TServerConfig
	(*TNullServiceConfig)(nil),   // 1: NCloud.NFileStore.NProto.TNullServiceConfig
	(*TLocalServiceConfig)(nil),  // 2: NCloud.NFileStore.NProto.TLocalServiceConfig
	(*TKikimrServiceConfig)(nil), // 3: NCloud.NFileStore.NProto.TKikimrServiceConfig
	(*TServerAppConfig)(nil),     // 4: NCloud.NFileStore.NProto.TServerAppConfig
	(*protos.TCertificate)(nil),  // 5: NCloud.NProto.TCertificate
}
var file_cloud_filestore_config_server_proto_depIdxs = []int32{
	5, // 0: NCloud.NFileStore.NProto.TServerConfig.Certs:type_name -> NCloud.NProto.TCertificate
	0, // 1: NCloud.NFileStore.NProto.TServerAppConfig.ServerConfig:type_name -> NCloud.NFileStore.NProto.TServerConfig
	1, // 2: NCloud.NFileStore.NProto.TServerAppConfig.NullServiceConfig:type_name -> NCloud.NFileStore.NProto.TNullServiceConfig
	2, // 3: NCloud.NFileStore.NProto.TServerAppConfig.LocalServiceConfig:type_name -> NCloud.NFileStore.NProto.TLocalServiceConfig
	3, // 4: NCloud.NFileStore.NProto.TServerAppConfig.KikimrServiceConfig:type_name -> NCloud.NFileStore.NProto.TKikimrServiceConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cloud_filestore_config_server_proto_init() }
func file_cloud_filestore_config_server_proto_init() {
	if File_cloud_filestore_config_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_filestore_config_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TServerConfig); i {
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
		file_cloud_filestore_config_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TNullServiceConfig); i {
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
		file_cloud_filestore_config_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLocalServiceConfig); i {
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
		file_cloud_filestore_config_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TKikimrServiceConfig); i {
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
		file_cloud_filestore_config_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TServerAppConfig); i {
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
			RawDescriptor: file_cloud_filestore_config_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_filestore_config_server_proto_goTypes,
		DependencyIndexes: file_cloud_filestore_config_server_proto_depIdxs,
		MessageInfos:      file_cloud_filestore_config_server_proto_msgTypes,
	}.Build()
	File_cloud_filestore_config_server_proto = out.File
	file_cloud_filestore_config_server_proto_rawDesc = nil
	file_cloud_filestore_config_server_proto_goTypes = nil
	file_cloud_filestore_config_server_proto_depIdxs = nil
}
