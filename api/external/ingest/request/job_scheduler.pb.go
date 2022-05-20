// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/ingest/request/job_scheduler.proto

package request

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

// SchedulerConfig
// The job message to configure the Delete Node Job
// every - It accepts '1h30m', '1m', '2h30m', ...
type SchedulerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Every     string `protobuf:"bytes,1,opt,name=every,proto3" json:"every,omitempty"`
	Threshold string `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Running   bool   `protobuf:"varint,3,opt,name=running,proto3" json:"running,omitempty"`
}

func (x *SchedulerConfig) Reset() {
	*x = SchedulerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerConfig) ProtoMessage() {}

func (x *SchedulerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerConfig.ProtoReflect.Descriptor instead.
func (*SchedulerConfig) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulerConfig) GetEvery() string {
	if x != nil {
		return x.Every
	}
	return ""
}

func (x *SchedulerConfig) GetThreshold() string {
	if x != nil {
		return x.Threshold
	}
	return ""
}

func (x *SchedulerConfig) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

type GetStatusJobScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusJobScheduler) Reset() {
	*x = GetStatusJobScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusJobScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusJobScheduler) ProtoMessage() {}

func (x *GetStatusJobScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusJobScheduler.ProtoReflect.Descriptor instead.
func (*GetStatusJobScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{1}
}

// DeleteNodesScheduler
type DeleteMarkedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMarkedNodes) Reset() {
	*x = DeleteMarkedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMarkedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMarkedNodes) ProtoMessage() {}

func (x *DeleteMarkedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMarkedNodes.ProtoReflect.Descriptor instead.
func (*DeleteMarkedNodes) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{2}
}

type StartDeleteNodesScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartDeleteNodesScheduler) Reset() {
	*x = StartDeleteNodesScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDeleteNodesScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDeleteNodesScheduler) ProtoMessage() {}

func (x *StartDeleteNodesScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDeleteNodesScheduler.ProtoReflect.Descriptor instead.
func (*StartDeleteNodesScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{3}
}

type StopDeleteNodesScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopDeleteNodesScheduler) Reset() {
	*x = StopDeleteNodesScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDeleteNodesScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDeleteNodesScheduler) ProtoMessage() {}

func (x *StopDeleteNodesScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDeleteNodesScheduler.ProtoReflect.Descriptor instead.
func (*StopDeleteNodesScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{4}
}

// MissingNodesForDeletionScheduler
type MarkMissingNodesForDeletion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkMissingNodesForDeletion) Reset() {
	*x = MarkMissingNodesForDeletion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkMissingNodesForDeletion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMissingNodesForDeletion) ProtoMessage() {}

func (x *MarkMissingNodesForDeletion) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMissingNodesForDeletion.ProtoReflect.Descriptor instead.
func (*MarkMissingNodesForDeletion) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{5}
}

type StartMissingNodesForDeletionScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartMissingNodesForDeletionScheduler) Reset() {
	*x = StartMissingNodesForDeletionScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMissingNodesForDeletionScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMissingNodesForDeletionScheduler) ProtoMessage() {}

func (x *StartMissingNodesForDeletionScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMissingNodesForDeletionScheduler.ProtoReflect.Descriptor instead.
func (*StartMissingNodesForDeletionScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{6}
}

type StopMissingNodesForDeletionScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopMissingNodesForDeletionScheduler) Reset() {
	*x = StopMissingNodesForDeletionScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopMissingNodesForDeletionScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopMissingNodesForDeletionScheduler) ProtoMessage() {}

func (x *StopMissingNodesForDeletionScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopMissingNodesForDeletionScheduler.ProtoReflect.Descriptor instead.
func (*StopMissingNodesForDeletionScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{7}
}

// NodesMissingScheduler
type MarkNodesMissing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkNodesMissing) Reset() {
	*x = MarkNodesMissing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNodesMissing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNodesMissing) ProtoMessage() {}

func (x *MarkNodesMissing) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNodesMissing.ProtoReflect.Descriptor instead.
func (*MarkNodesMissing) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{8}
}

type StartNodesMissingScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartNodesMissingScheduler) Reset() {
	*x = StartNodesMissingScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartNodesMissingScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartNodesMissingScheduler) ProtoMessage() {}

func (x *StartNodesMissingScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartNodesMissingScheduler.ProtoReflect.Descriptor instead.
func (*StartNodesMissingScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{9}
}

type StopNodesMissingScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopNodesMissingScheduler) Reset() {
	*x = StopNodesMissingScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopNodesMissingScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopNodesMissingScheduler) ProtoMessage() {}

func (x *StopNodesMissingScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_job_scheduler_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopNodesMissingScheduler.ProtoReflect.Descriptor instead.
func (*StopNodesMissingScheduler) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_job_scheduler_proto_rawDescGZIP(), []int{10}
}

var File_external_ingest_request_job_scheduler_proto protoreflect.FileDescriptor

var file_external_ingest_request_job_scheduler_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x5f, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f, 0x62,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x1b,
	0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x18, 0x53,
	0x74, 0x6f, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x61, 0x72, 0x6b, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x25, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22,
	0x26, 0x0a, 0x24, 0x53, 0x74, 0x6f, 0x70, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x12, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x1c, 0x0a, 0x1a, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x74, 0x6f,
	0x70, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_ingest_request_job_scheduler_proto_rawDescOnce sync.Once
	file_external_ingest_request_job_scheduler_proto_rawDescData = file_external_ingest_request_job_scheduler_proto_rawDesc
)

func file_external_ingest_request_job_scheduler_proto_rawDescGZIP() []byte {
	file_external_ingest_request_job_scheduler_proto_rawDescOnce.Do(func() {
		file_external_ingest_request_job_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_ingest_request_job_scheduler_proto_rawDescData)
	})
	return file_external_ingest_request_job_scheduler_proto_rawDescData
}

var file_external_ingest_request_job_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_external_ingest_request_job_scheduler_proto_goTypes = []interface{}{
	(*SchedulerConfig)(nil),                       // 0: chef.automate.api.ingest.request.SchedulerConfig
	(*GetStatusJobScheduler)(nil),                 // 1: chef.automate.api.ingest.request.GetStatusJobScheduler
	(*DeleteMarkedNodes)(nil),                     // 2: chef.automate.api.ingest.request.DeleteMarkedNodes
	(*StartDeleteNodesScheduler)(nil),             // 3: chef.automate.api.ingest.request.StartDeleteNodesScheduler
	(*StopDeleteNodesScheduler)(nil),              // 4: chef.automate.api.ingest.request.StopDeleteNodesScheduler
	(*MarkMissingNodesForDeletion)(nil),           // 5: chef.automate.api.ingest.request.MarkMissingNodesForDeletion
	(*StartMissingNodesForDeletionScheduler)(nil), // 6: chef.automate.api.ingest.request.StartMissingNodesForDeletionScheduler
	(*StopMissingNodesForDeletionScheduler)(nil),  // 7: chef.automate.api.ingest.request.StopMissingNodesForDeletionScheduler
	(*MarkNodesMissing)(nil),                      // 8: chef.automate.api.ingest.request.MarkNodesMissing
	(*StartNodesMissingScheduler)(nil),            // 9: chef.automate.api.ingest.request.StartNodesMissingScheduler
	(*StopNodesMissingScheduler)(nil),             // 10: chef.automate.api.ingest.request.StopNodesMissingScheduler
}
var file_external_ingest_request_job_scheduler_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_ingest_request_job_scheduler_proto_init() }
func file_external_ingest_request_job_scheduler_proto_init() {
	if File_external_ingest_request_job_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_ingest_request_job_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerConfig); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusJobScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMarkedNodes); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDeleteNodesScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopDeleteNodesScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkMissingNodesForDeletion); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMissingNodesForDeletionScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopMissingNodesForDeletionScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkNodesMissing); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartNodesMissingScheduler); i {
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
		file_external_ingest_request_job_scheduler_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopNodesMissingScheduler); i {
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
			RawDescriptor: file_external_ingest_request_job_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_ingest_request_job_scheduler_proto_goTypes,
		DependencyIndexes: file_external_ingest_request_job_scheduler_proto_depIdxs,
		MessageInfos:      file_external_ingest_request_job_scheduler_proto_msgTypes,
	}.Build()
	File_external_ingest_request_job_scheduler_proto = out.File
	file_external_ingest_request_job_scheduler_proto_rawDesc = nil
	file_external_ingest_request_job_scheduler_proto_goTypes = nil
	file_external_ingest_request_job_scheduler_proto_depIdxs = nil
}
