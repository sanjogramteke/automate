// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: config/id/config_request.proto

package id

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if x != nil {
		return x.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa    *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Log     *ConfigRequest_V1_System_Log     `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Network *ConfigRequest_V1_System_Network `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty" toml:"network,omitempty" mapstructure:"network,omitempty"`
	Sql     *ConfigRequest_V1_System_Sql     `protobuf:"bytes,4,opt,name=sql,proto3" json:"sql,omitempty" toml:"sql,omitempty" mapstructure:"sql,omitempty"`
	Tls     *shared.TLSCredentials           `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetNetwork() *ConfigRequest_V1_System_Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetSql() *ConfigRequest_V1_System_Sql {
	if x != nil {
		return x.Sql
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest_V1_Service) Reset() {
	*x = ConfigRequest_V1_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0, 1}
}

type ConfigRequest_V1_System_Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	// Deprecated: Do not use.
	ListenIp *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=listen_ip,json=listenIp,proto3" json:"listen_ip,omitempty" toml:"listen_ip,omitempty" mapstructure:"listen_ip,omitempty"` // The listen host is no longer setable(localhost only)
}

func (x *ConfigRequest_V1_System_Network) Reset() {
	*x = ConfigRequest_V1_System_Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Network) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Network) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Network.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Network) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ConfigRequest_V1_System_Network) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Network) GetListenIp() *wrapperspb.StringValue {
	if x != nil {
		return x.ListenIp
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level                 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	RotationMaxBytes      *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=rotation_max_bytes,json=rotationMaxBytes,proto3" json:"rotation_max_bytes,omitempty" toml:"rotation_max_bytes,omitempty" mapstructure:"rotation_max_bytes,omitempty"`
	RotationMaxFiles      *wrapperspb.Int32Value  `protobuf:"bytes,3,opt,name=rotation_max_files,json=rotationMaxFiles,proto3" json:"rotation_max_files,omitempty" toml:"rotation_max_files,omitempty" mapstructure:"rotation_max_files,omitempty"`
	MaxErrorLogsPerSecond *wrapperspb.Int32Value  `protobuf:"bytes,4,opt,name=max_error_logs_per_second,json=maxErrorLogsPerSecond,proto3" json:"max_error_logs_per_second,omitempty" toml:"max_error_logs_per_second,omitempty" mapstructure:"max_error_logs_per_second,omitempty"`
	// TODO(ssd) 2018-07-24: Different log
	// rotation configurables require
	// different units.
	RotationMaxMegabytes *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=rotation_max_megabytes,json=rotationMaxMegabytes,proto3" json:"rotation_max_megabytes,omitempty" toml:"rotation_max_megabytes,omitempty" mapstructure:"rotation_max_megabytes,omitempty"`
	ExtendedPerfLog      *wrapperspb.BoolValue  `protobuf:"bytes,6,opt,name=extended_perf_log,json=extendedPerfLog,proto3" json:"extended_perf_log,omitempty" toml:"extended_perf_log,omitempty" mapstructure:"extended_perf_log,omitempty"`
}

func (x *ConfigRequest_V1_System_Log) Reset() {
	*x = ConfigRequest_V1_System_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Log) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Log) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Log.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *ConfigRequest_V1_System_Log) GetLevel() *wrapperspb.StringValue {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetRotationMaxBytes() *wrapperspb.Int64Value {
	if x != nil {
		return x.RotationMaxBytes
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetRotationMaxFiles() *wrapperspb.Int32Value {
	if x != nil {
		return x.RotationMaxFiles
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetMaxErrorLogsPerSecond() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxErrorLogsPerSecond
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetRotationMaxMegabytes() *wrapperspb.Int32Value {
	if x != nil {
		return x.RotationMaxMegabytes
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetExtendedPerfLog() *wrapperspb.BoolValue {
	if x != nil {
		return x.ExtendedPerfLog
	}
	return nil
}

type ConfigRequest_V1_System_Sql struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout          *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty" toml:"timeout,omitempty" mapstructure:"timeout,omitempty"`
	PoolInitSize     *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=pool_init_size,json=poolInitSize,proto3" json:"pool_init_size,omitempty" toml:"pool_init_size,omitempty" mapstructure:"pool_init_size,omitempty"`
	PoolMaxSize      *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=pool_max_size,json=poolMaxSize,proto3" json:"pool_max_size,omitempty" toml:"pool_max_size,omitempty" mapstructure:"pool_max_size,omitempty"`
	PoolQueueMax     *wrapperspb.Int32Value `protobuf:"bytes,4,opt,name=pool_queue_max,json=poolQueueMax,proto3" json:"pool_queue_max,omitempty" toml:"pool_queue_max,omitempty" mapstructure:"pool_queue_max,omitempty"`
	PoolQueueTimeout *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=pool_queue_timeout,json=poolQueueTimeout,proto3" json:"pool_queue_timeout,omitempty" toml:"pool_queue_timeout,omitempty" mapstructure:"pool_queue_timeout,omitempty"`
}

func (x *ConfigRequest_V1_System_Sql) Reset() {
	*x = ConfigRequest_V1_System_Sql{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_id_config_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Sql) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Sql) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Sql) ProtoReflect() protoreflect.Message {
	mi := &file_config_id_config_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Sql.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Sql) Descriptor() ([]byte, []int) {
	return file_config_id_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 2}
}

func (x *ConfigRequest_V1_System_Sql) GetTimeout() *wrapperspb.Int32Value {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *ConfigRequest_V1_System_Sql) GetPoolInitSize() *wrapperspb.Int32Value {
	if x != nil {
		return x.PoolInitSize
	}
	return nil
}

func (x *ConfigRequest_V1_System_Sql) GetPoolMaxSize() *wrapperspb.Int32Value {
	if x != nil {
		return x.PoolMaxSize
	}
	return nil
}

func (x *ConfigRequest_V1_System_Sql) GetPoolQueueMax() *wrapperspb.Int32Value {
	if x != nil {
		return x.PoolQueueMax
	}
	return nil
}

func (x *ConfigRequest_V1_System_Sql) GetPoolQueueTimeout() *wrapperspb.Int32Value {
	if x != nil {
		return x.PoolQueueTimeout
	}
	return nil
}

var File_config_id_config_request_proto protoreflect.FileDescriptor

var file_config_id_config_request_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x66, 0x72,
	0x6f, 0x73, 0x74, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x61, 0x32, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x32, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x0c, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x02, 0x76,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31, 0x1a, 0xbe, 0x0b, 0x0a, 0x02,
	0x56, 0x31, 0x12, 0x47, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x48, 0x0a, 0x03, 0x73,
	0x76, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x03, 0x73, 0x76, 0x63, 0x1a, 0x99, 0x0a, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6c, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6c, 0x73, 0x61,
	0x52, 0x04, 0x6d, 0x6c, 0x73, 0x61, 0x12, 0x4b, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x66, 0x72, 0x6f,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x12, 0x57, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x66, 0x72,
	0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x4b, 0x0a, 0x03,
	0x73, 0x71, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x53, 0x71, 0x6c, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x1a, 0x92, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x48, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x17,
	0xc2, 0xf3, 0x18, 0x13, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0xda, 0x4f,
	0x1a, 0x05, 0x68, 0x74, 0x74, 0x70, 0x73, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x49, 0x70, 0x1a, 0xc1, 0x03, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x49, 0x0a, 0x12, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x10, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x12, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x55,
	0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15,
	0x6d, 0x61, 0x78, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x51, 0x0a, 0x16, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x14, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x4d,
	0x65, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x4c, 0x6f, 0x67,
	0x1a, 0xce, 0x02, 0x0a, 0x03, 0x53, 0x71, 0x6c, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x41, 0x0a, 0x0e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x78, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x49, 0x0a, 0x12, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x70, 0x6f, 0x6f, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x1a, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x1c, 0xc2, 0xf3,
	0x18, 0x18, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x73, 0x2d,
	0x6f, 0x63, 0x2d, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_id_config_request_proto_rawDescOnce sync.Once
	file_config_id_config_request_proto_rawDescData = file_config_id_config_request_proto_rawDesc
)

func file_config_id_config_request_proto_rawDescGZIP() []byte {
	file_config_id_config_request_proto_rawDescOnce.Do(func() {
		file_config_id_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_id_config_request_proto_rawDescData)
	})
	return file_config_id_config_request_proto_rawDescData
}

var file_config_id_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_id_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                   // 0: chef.automate.domain.id.ConfigRequest
	(*ConfigRequest_V1)(nil),                // 1: chef.automate.domain.id.ConfigRequest.V1
	(*ConfigRequest_V1_System)(nil),         // 2: chef.automate.domain.id.ConfigRequest.V1.System
	(*ConfigRequest_V1_Service)(nil),        // 3: chef.automate.domain.id.ConfigRequest.V1.Service
	(*ConfigRequest_V1_System_Network)(nil), // 4: chef.automate.domain.id.ConfigRequest.V1.System.Network
	(*ConfigRequest_V1_System_Log)(nil),     // 5: chef.automate.domain.id.ConfigRequest.V1.System.Log
	(*ConfigRequest_V1_System_Sql)(nil),     // 6: chef.automate.domain.id.ConfigRequest.V1.System.Sql
	(*shared.Mlsa)(nil),                     // 7: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),           // 8: chef.automate.infra.config.TLSCredentials
	(*wrapperspb.Int32Value)(nil),           // 9: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil),          // 10: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),           // 11: google.protobuf.Int64Value
	(*wrapperspb.BoolValue)(nil),            // 12: google.protobuf.BoolValue
}
var file_config_id_config_request_proto_depIdxs = []int32{
	1,  // 0: chef.automate.domain.id.ConfigRequest.v1:type_name -> chef.automate.domain.id.ConfigRequest.V1
	2,  // 1: chef.automate.domain.id.ConfigRequest.V1.sys:type_name -> chef.automate.domain.id.ConfigRequest.V1.System
	3,  // 2: chef.automate.domain.id.ConfigRequest.V1.svc:type_name -> chef.automate.domain.id.ConfigRequest.V1.Service
	7,  // 3: chef.automate.domain.id.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	5,  // 4: chef.automate.domain.id.ConfigRequest.V1.System.log:type_name -> chef.automate.domain.id.ConfigRequest.V1.System.Log
	4,  // 5: chef.automate.domain.id.ConfigRequest.V1.System.network:type_name -> chef.automate.domain.id.ConfigRequest.V1.System.Network
	6,  // 6: chef.automate.domain.id.ConfigRequest.V1.System.sql:type_name -> chef.automate.domain.id.ConfigRequest.V1.System.Sql
	8,  // 7: chef.automate.domain.id.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	9,  // 8: chef.automate.domain.id.ConfigRequest.V1.System.Network.port:type_name -> google.protobuf.Int32Value
	10, // 9: chef.automate.domain.id.ConfigRequest.V1.System.Network.listen_ip:type_name -> google.protobuf.StringValue
	10, // 10: chef.automate.domain.id.ConfigRequest.V1.System.Log.level:type_name -> google.protobuf.StringValue
	11, // 11: chef.automate.domain.id.ConfigRequest.V1.System.Log.rotation_max_bytes:type_name -> google.protobuf.Int64Value
	9,  // 12: chef.automate.domain.id.ConfigRequest.V1.System.Log.rotation_max_files:type_name -> google.protobuf.Int32Value
	9,  // 13: chef.automate.domain.id.ConfigRequest.V1.System.Log.max_error_logs_per_second:type_name -> google.protobuf.Int32Value
	9,  // 14: chef.automate.domain.id.ConfigRequest.V1.System.Log.rotation_max_megabytes:type_name -> google.protobuf.Int32Value
	12, // 15: chef.automate.domain.id.ConfigRequest.V1.System.Log.extended_perf_log:type_name -> google.protobuf.BoolValue
	9,  // 16: chef.automate.domain.id.ConfigRequest.V1.System.Sql.timeout:type_name -> google.protobuf.Int32Value
	9,  // 17: chef.automate.domain.id.ConfigRequest.V1.System.Sql.pool_init_size:type_name -> google.protobuf.Int32Value
	9,  // 18: chef.automate.domain.id.ConfigRequest.V1.System.Sql.pool_max_size:type_name -> google.protobuf.Int32Value
	9,  // 19: chef.automate.domain.id.ConfigRequest.V1.System.Sql.pool_queue_max:type_name -> google.protobuf.Int32Value
	9,  // 20: chef.automate.domain.id.ConfigRequest.V1.System.Sql.pool_queue_timeout:type_name -> google.protobuf.Int32Value
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_config_id_config_request_proto_init() }
func file_config_id_config_request_proto_init() {
	if File_config_id_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_id_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_id_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
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
		file_config_id_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
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
		file_config_id_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_Service); i {
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
		file_config_id_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Network); i {
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
		file_config_id_config_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Log); i {
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
		file_config_id_config_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Sql); i {
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
			RawDescriptor: file_config_id_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_id_config_request_proto_goTypes,
		DependencyIndexes: file_config_id_config_request_proto_depIdxs,
		MessageInfos:      file_config_id_config_request_proto_msgTypes,
	}.Build()
	File_config_id_config_request_proto = out.File
	file_config_id_config_request_proto_rawDesc = nil
	file_config_id_config_request_proto_goTypes = nil
	file_config_id_config_request_proto_depIdxs = nil
}
