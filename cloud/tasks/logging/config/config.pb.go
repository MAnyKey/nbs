// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/tasks/logging/config/config.proto

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

type Level int32

const (
	Level_LEVEL_TRACE Level = 0
	Level_LEVEL_DEBUG Level = 1
	Level_LEVEL_INFO  Level = 2
	Level_LEVEL_WARN  Level = 3
	Level_LEVEL_ERROR Level = 4
	Level_LEVEL_FATAL Level = 5
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "LEVEL_TRACE",
		1: "LEVEL_DEBUG",
		2: "LEVEL_INFO",
		3: "LEVEL_WARN",
		4: "LEVEL_ERROR",
		5: "LEVEL_FATAL",
	}
	Level_value = map[string]int32{
		"LEVEL_TRACE": 0,
		"LEVEL_DEBUG": 1,
		"LEVEL_INFO":  2,
		"LEVEL_WARN":  3,
		"LEVEL_ERROR": 4,
		"LEVEL_FATAL": 5,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_tasks_logging_config_config_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_cloud_tasks_logging_config_config_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Level) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Level(num)
	return nil
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_cloud_tasks_logging_config_config_proto_rawDescGZIP(), []int{0}
}

type StderrLoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StderrLoggingConfig) Reset() {
	*x = StderrLoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StderrLoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StderrLoggingConfig) ProtoMessage() {}

func (x *StderrLoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StderrLoggingConfig.ProtoReflect.Descriptor instead.
func (*StderrLoggingConfig) Descriptor() ([]byte, []int) {
	return file_cloud_tasks_logging_config_config_proto_rawDescGZIP(), []int{0}
}

type JournaldLoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JournaldLoggingConfig) Reset() {
	*x = JournaldLoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournaldLoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournaldLoggingConfig) ProtoMessage() {}

func (x *JournaldLoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournaldLoggingConfig.ProtoReflect.Descriptor instead.
func (*JournaldLoggingConfig) Descriptor() ([]byte, []int) {
	return file_cloud_tasks_logging_config_config_proto_rawDescGZIP(), []int{1}
}

type LoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Logging:
	//
	//	*LoggingConfig_LoggingStderr
	//	*LoggingConfig_LoggingJournald
	Logging isLoggingConfig_Logging `protobuf_oneof:"Logging"`
	Level   *Level                  `protobuf:"varint,3,opt,name=Level,enum=logging.Level,def=2" json:"Level,omitempty"`
}

// Default values for LoggingConfig fields.
const (
	Default_LoggingConfig_Level = Level_LEVEL_INFO
)

func (x *LoggingConfig) Reset() {
	*x = LoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingConfig) ProtoMessage() {}

func (x *LoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_tasks_logging_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingConfig.ProtoReflect.Descriptor instead.
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return file_cloud_tasks_logging_config_config_proto_rawDescGZIP(), []int{2}
}

func (m *LoggingConfig) GetLogging() isLoggingConfig_Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (x *LoggingConfig) GetLoggingStderr() *StderrLoggingConfig {
	if x, ok := x.GetLogging().(*LoggingConfig_LoggingStderr); ok {
		return x.LoggingStderr
	}
	return nil
}

func (x *LoggingConfig) GetLoggingJournald() *JournaldLoggingConfig {
	if x, ok := x.GetLogging().(*LoggingConfig_LoggingJournald); ok {
		return x.LoggingJournald
	}
	return nil
}

func (x *LoggingConfig) GetLevel() Level {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return Default_LoggingConfig_Level
}

type isLoggingConfig_Logging interface {
	isLoggingConfig_Logging()
}

type LoggingConfig_LoggingStderr struct {
	LoggingStderr *StderrLoggingConfig `protobuf:"bytes,1,opt,name=LoggingStderr,oneof"`
}

type LoggingConfig_LoggingJournald struct {
	LoggingJournald *JournaldLoggingConfig `protobuf:"bytes,2,opt,name=LoggingJournald,oneof"`
}

func (*LoggingConfig_LoggingStderr) isLoggingConfig_Logging() {}

func (*LoggingConfig_LoggingJournald) isLoggingConfig_Logging() {}

var File_cloud_tasks_logging_config_config_proto protoreflect.FileDescriptor

var file_cloud_tasks_logging_config_config_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x17, 0x0a, 0x15, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xde, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0d, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x4a, 0x0a, 0x0f, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2a, 0x6b, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0f, 0x0a, 0x0b,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x05,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x62, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
}

var (
	file_cloud_tasks_logging_config_config_proto_rawDescOnce sync.Once
	file_cloud_tasks_logging_config_config_proto_rawDescData = file_cloud_tasks_logging_config_config_proto_rawDesc
)

func file_cloud_tasks_logging_config_config_proto_rawDescGZIP() []byte {
	file_cloud_tasks_logging_config_config_proto_rawDescOnce.Do(func() {
		file_cloud_tasks_logging_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_tasks_logging_config_config_proto_rawDescData)
	})
	return file_cloud_tasks_logging_config_config_proto_rawDescData
}

var file_cloud_tasks_logging_config_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_tasks_logging_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_tasks_logging_config_config_proto_goTypes = []interface{}{
	(Level)(0),                    // 0: logging.Level
	(*StderrLoggingConfig)(nil),   // 1: logging.StderrLoggingConfig
	(*JournaldLoggingConfig)(nil), // 2: logging.JournaldLoggingConfig
	(*LoggingConfig)(nil),         // 3: logging.LoggingConfig
}
var file_cloud_tasks_logging_config_config_proto_depIdxs = []int32{
	1, // 0: logging.LoggingConfig.LoggingStderr:type_name -> logging.StderrLoggingConfig
	2, // 1: logging.LoggingConfig.LoggingJournald:type_name -> logging.JournaldLoggingConfig
	0, // 2: logging.LoggingConfig.Level:type_name -> logging.Level
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_tasks_logging_config_config_proto_init() }
func file_cloud_tasks_logging_config_config_proto_init() {
	if File_cloud_tasks_logging_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_tasks_logging_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StderrLoggingConfig); i {
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
		file_cloud_tasks_logging_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournaldLoggingConfig); i {
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
		file_cloud_tasks_logging_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingConfig); i {
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
	file_cloud_tasks_logging_config_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*LoggingConfig_LoggingStderr)(nil),
		(*LoggingConfig_LoggingJournald)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_tasks_logging_config_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_tasks_logging_config_config_proto_goTypes,
		DependencyIndexes: file_cloud_tasks_logging_config_config_proto_depIdxs,
		EnumInfos:         file_cloud_tasks_logging_config_config_proto_enumTypes,
		MessageInfos:      file_cloud_tasks_logging_config_config_proto_msgTypes,
	}.Build()
	File_cloud_tasks_logging_config_config_proto = out.File
	file_cloud_tasks_logging_config_config_proto_rawDesc = nil
	file_cloud_tasks_logging_config_config_proto_goTypes = nil
	file_cloud_tasks_logging_config_config_proto_depIdxs = nil
}
