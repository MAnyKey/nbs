// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/filestore/config/client.proto

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

type TClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                  *string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	Port                  *uint32 `protobuf:"varint,2,opt,name=Port" json:"Port,omitempty"`
	RetryTimeout          *uint32 `protobuf:"varint,3,opt,name=RetryTimeout" json:"RetryTimeout,omitempty"`
	RetryTimeoutIncrement *uint32 `protobuf:"varint,4,opt,name=RetryTimeoutIncrement" json:"RetryTimeoutIncrement,omitempty"`
	ThreadsCount          *uint32 `protobuf:"varint,5,opt,name=ThreadsCount" json:"ThreadsCount,omitempty"`
	// Request timeout (in milliseconds).
	RequestTimeout *uint32 `protobuf:"varint,6,opt,name=RequestTimeout" json:"RequestTimeout,omitempty"`
	// Max retry timeout in case of connection failure (in milliseconds).
	ConnectionErrorMaxRetryTimeout *uint32 `protobuf:"varint,7,opt,name=ConnectionErrorMaxRetryTimeout" json:"ConnectionErrorMaxRetryTimeout,omitempty"`
	// Max message size for sending/receiving in bytes.
	MaxMessageSize       *uint32 `protobuf:"varint,8,opt,name=MaxMessageSize" json:"MaxMessageSize,omitempty"`
	MemoryQuotaBytes     *uint32 `protobuf:"varint,9,opt,name=MemoryQuotaBytes" json:"MemoryQuotaBytes,omitempty"`
	GrpcThreadsLimit     *uint32 `protobuf:"varint,10,opt,name=GrpcThreadsLimit" json:"GrpcThreadsLimit,omitempty"`
	GrpcReconnectBackoff *uint32 `protobuf:"varint,11,opt,name=GrpcReconnectBackoff" json:"GrpcReconnectBackoff,omitempty"`
	// Remote host secure control port (with TLS).
	// If set, client will ignore Port and connect via SecurePort.
	SecurePort *uint32 `protobuf:"varint,12,opt,name=SecurePort" json:"SecurePort,omitempty"`
	// TLS details.
	RootCertsFile      *string `protobuf:"bytes,13,opt,name=RootCertsFile" json:"RootCertsFile,omitempty"`
	CertFile           *string `protobuf:"bytes,14,opt,name=CertFile" json:"CertFile,omitempty"`
	CertPrivateKeyFile *string `protobuf:"bytes,15,opt,name=CertPrivateKeyFile" json:"CertPrivateKeyFile,omitempty"`
	// Token to use for authentication.
	AuthToken *string `protobuf:"bytes,16,opt,name=AuthToken" json:"AuthToken,omitempty"`
	// Skip server certificate verification
	SkipCertVerification *bool `protobuf:"varint,17,opt,name=SkipCertVerification" json:"SkipCertVerification,omitempty"`
	// Unix-socket path.
	UnixSocketPath *string `protobuf:"bytes,18,opt,name=UnixSocketPath" json:"UnixSocketPath,omitempty"`
}

func (x *TClientConfig) Reset() {
	*x = TClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientConfig) ProtoMessage() {}

func (x *TClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientConfig.ProtoReflect.Descriptor instead.
func (*TClientConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_client_proto_rawDescGZIP(), []int{0}
}

func (x *TClientConfig) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *TClientConfig) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *TClientConfig) GetRetryTimeout() uint32 {
	if x != nil && x.RetryTimeout != nil {
		return *x.RetryTimeout
	}
	return 0
}

func (x *TClientConfig) GetRetryTimeoutIncrement() uint32 {
	if x != nil && x.RetryTimeoutIncrement != nil {
		return *x.RetryTimeoutIncrement
	}
	return 0
}

func (x *TClientConfig) GetThreadsCount() uint32 {
	if x != nil && x.ThreadsCount != nil {
		return *x.ThreadsCount
	}
	return 0
}

func (x *TClientConfig) GetRequestTimeout() uint32 {
	if x != nil && x.RequestTimeout != nil {
		return *x.RequestTimeout
	}
	return 0
}

func (x *TClientConfig) GetConnectionErrorMaxRetryTimeout() uint32 {
	if x != nil && x.ConnectionErrorMaxRetryTimeout != nil {
		return *x.ConnectionErrorMaxRetryTimeout
	}
	return 0
}

func (x *TClientConfig) GetMaxMessageSize() uint32 {
	if x != nil && x.MaxMessageSize != nil {
		return *x.MaxMessageSize
	}
	return 0
}

func (x *TClientConfig) GetMemoryQuotaBytes() uint32 {
	if x != nil && x.MemoryQuotaBytes != nil {
		return *x.MemoryQuotaBytes
	}
	return 0
}

func (x *TClientConfig) GetGrpcThreadsLimit() uint32 {
	if x != nil && x.GrpcThreadsLimit != nil {
		return *x.GrpcThreadsLimit
	}
	return 0
}

func (x *TClientConfig) GetGrpcReconnectBackoff() uint32 {
	if x != nil && x.GrpcReconnectBackoff != nil {
		return *x.GrpcReconnectBackoff
	}
	return 0
}

func (x *TClientConfig) GetSecurePort() uint32 {
	if x != nil && x.SecurePort != nil {
		return *x.SecurePort
	}
	return 0
}

func (x *TClientConfig) GetRootCertsFile() string {
	if x != nil && x.RootCertsFile != nil {
		return *x.RootCertsFile
	}
	return ""
}

func (x *TClientConfig) GetCertFile() string {
	if x != nil && x.CertFile != nil {
		return *x.CertFile
	}
	return ""
}

func (x *TClientConfig) GetCertPrivateKeyFile() string {
	if x != nil && x.CertPrivateKeyFile != nil {
		return *x.CertPrivateKeyFile
	}
	return ""
}

func (x *TClientConfig) GetAuthToken() string {
	if x != nil && x.AuthToken != nil {
		return *x.AuthToken
	}
	return ""
}

func (x *TClientConfig) GetSkipCertVerification() bool {
	if x != nil && x.SkipCertVerification != nil {
		return *x.SkipCertVerification
	}
	return false
}

func (x *TClientConfig) GetUnixSocketPath() string {
	if x != nil && x.UnixSocketPath != nil {
		return *x.UnixSocketPath
	}
	return ""
}

type TClientAppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientConfig *TClientConfig `protobuf:"bytes,1,opt,name=ClientConfig" json:"ClientConfig,omitempty"`
}

func (x *TClientAppConfig) Reset() {
	*x = TClientAppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TClientAppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientAppConfig) ProtoMessage() {}

func (x *TClientAppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientAppConfig.ProtoReflect.Descriptor instead.
func (*TClientAppConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_client_proto_rawDescGZIP(), []int{1}
}

func (x *TClientAppConfig) GetClientConfig() *TClientConfig {
	if x != nil {
		return x.ClientConfig
	}
	return nil
}

type TSessionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSystemId        *string `protobuf:"bytes,1,opt,name=FileSystemId" json:"FileSystemId,omitempty"`
	ClientId            *string `protobuf:"bytes,2,opt,name=ClientId" json:"ClientId,omitempty"`
	SessionPingTimeout  *uint32 `protobuf:"varint,3,opt,name=SessionPingTimeout" json:"SessionPingTimeout,omitempty"`
	SessionRetryTimeout *uint32 `protobuf:"varint,4,opt,name=SessionRetryTimeout" json:"SessionRetryTimeout,omitempty"`
}

func (x *TSessionConfig) Reset() {
	*x = TSessionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_filestore_config_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TSessionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TSessionConfig) ProtoMessage() {}

func (x *TSessionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_filestore_config_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TSessionConfig.ProtoReflect.Descriptor instead.
func (*TSessionConfig) Descriptor() ([]byte, []int) {
	return file_cloud_filestore_config_client_proto_rawDescGZIP(), []int{2}
}

func (x *TSessionConfig) GetFileSystemId() string {
	if x != nil && x.FileSystemId != nil {
		return *x.FileSystemId
	}
	return ""
}

func (x *TSessionConfig) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *TSessionConfig) GetSessionPingTimeout() uint32 {
	if x != nil && x.SessionPingTimeout != nil {
		return *x.SessionPingTimeout
	}
	return 0
}

func (x *TSessionConfig) GetSessionRetryTimeout() uint32 {
	if x != nil && x.SessionRetryTimeout != nil {
		return *x.SessionRetryTimeout
	}
	return 0
}

var File_cloud_filestore_config_client_proto protoreflect.FileDescriptor

var file_cloud_filestore_config_client_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe5, 0x05, 0x0a, 0x0d, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x34, 0x0a,
	0x15, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x46, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x47,
	0x72, 0x70, 0x63, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x47, 0x72, 0x70, 0x63, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x52,
	0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a,
	0x12, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x65, 0x72, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x53,
	0x6b, 0x69, 0x70, 0x43, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x53, 0x6b, 0x69, 0x70, 0x43,
	0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x5f, 0x0a, 0x10, 0x54, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb2, 0x01, 0x0a, 0x0e, 0x54, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67,
}

var (
	file_cloud_filestore_config_client_proto_rawDescOnce sync.Once
	file_cloud_filestore_config_client_proto_rawDescData = file_cloud_filestore_config_client_proto_rawDesc
)

func file_cloud_filestore_config_client_proto_rawDescGZIP() []byte {
	file_cloud_filestore_config_client_proto_rawDescOnce.Do(func() {
		file_cloud_filestore_config_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_filestore_config_client_proto_rawDescData)
	})
	return file_cloud_filestore_config_client_proto_rawDescData
}

var file_cloud_filestore_config_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_filestore_config_client_proto_goTypes = []interface{}{
	(*TClientConfig)(nil),    // 0: NCloud.NFileStore.NProto.TClientConfig
	(*TClientAppConfig)(nil), // 1: NCloud.NFileStore.NProto.TClientAppConfig
	(*TSessionConfig)(nil),   // 2: NCloud.NFileStore.NProto.TSessionConfig
}
var file_cloud_filestore_config_client_proto_depIdxs = []int32{
	0, // 0: NCloud.NFileStore.NProto.TClientAppConfig.ClientConfig:type_name -> NCloud.NFileStore.NProto.TClientConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_filestore_config_client_proto_init() }
func file_cloud_filestore_config_client_proto_init() {
	if File_cloud_filestore_config_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_filestore_config_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TClientConfig); i {
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
		file_cloud_filestore_config_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TClientAppConfig); i {
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
		file_cloud_filestore_config_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TSessionConfig); i {
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
			RawDescriptor: file_cloud_filestore_config_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_filestore_config_client_proto_goTypes,
		DependencyIndexes: file_cloud_filestore_config_client_proto_depIdxs,
		MessageInfos:      file_cloud_filestore_config_client_proto_msgTypes,
	}.Build()
	File_cloud_filestore_config_client_proto = out.File
	file_cloud_filestore_config_client_proto_rawDesc = nil
	file_cloud_filestore_config_client_proto_goTypes = nil
	file_cloud_filestore_config_client_proto_depIdxs = nil
}
