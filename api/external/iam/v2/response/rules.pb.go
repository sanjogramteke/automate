// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/iam/v2/response/rules.proto

package response

import (
	common "github.com/chef/automate/api/external/iam/v2/common"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *common.Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *CreateRuleResp) Reset() {
	*x = CreateRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleResp) ProtoMessage() {}

func (x *CreateRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleResp.ProtoReflect.Descriptor instead.
func (*CreateRuleResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRuleResp) GetRule() *common.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type UpdateRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *common.Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *UpdateRuleResp) Reset() {
	*x = UpdateRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleResp) ProtoMessage() {}

func (x *UpdateRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleResp.ProtoReflect.Descriptor instead.
func (*UpdateRuleResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRuleResp) GetRule() *common.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type GetRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *common.Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *GetRuleResp) Reset() {
	*x = GetRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleResp) ProtoMessage() {}

func (x *GetRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleResp.ProtoReflect.Descriptor instead.
func (*GetRuleResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{2}
}

func (x *GetRuleResp) GetRule() *common.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type ListRulesForProjectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules  []*common.Rule            `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	Status common.ProjectRulesStatus `protobuf:"varint,2,opt,name=status,proto3,enum=chef.automate.api.iam.v2.ProjectRulesStatus" json:"status,omitempty"`
}

func (x *ListRulesForProjectResp) Reset() {
	*x = ListRulesForProjectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRulesForProjectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesForProjectResp) ProtoMessage() {}

func (x *ListRulesForProjectResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesForProjectResp.ProtoReflect.Descriptor instead.
func (*ListRulesForProjectResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{3}
}

func (x *ListRulesForProjectResp) GetRules() []*common.Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *ListRulesForProjectResp) GetStatus() common.ProjectRulesStatus {
	if x != nil {
		return x.Status
	}
	return common.ProjectRulesStatus(0)
}

type DeleteRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRuleResp) Reset() {
	*x = DeleteRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleResp) ProtoMessage() {}

func (x *DeleteRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleResp.ProtoReflect.Descriptor instead.
func (*DeleteRuleResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{4}
}

type ApplyRulesStartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyRulesStartResp) Reset() {
	*x = ApplyRulesStartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRulesStartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesStartResp) ProtoMessage() {}

func (x *ApplyRulesStartResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesStartResp.ProtoReflect.Descriptor instead.
func (*ApplyRulesStartResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{5}
}

type ApplyRulesCancelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyRulesCancelResp) Reset() {
	*x = ApplyRulesCancelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRulesCancelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesCancelResp) ProtoMessage() {}

func (x *ApplyRulesCancelResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesCancelResp.ProtoReflect.Descriptor instead.
func (*ApplyRulesCancelResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{6}
}

type ApplyRulesStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of two states: `not_running` and `running`.
	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// Estimated time when the project update will complete.
	EstimatedTimeComplete *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=estimated_time_complete,json=estimatedTimeComplete,proto3" json:"estimated_time_complete,omitempty"`
	// The percentage complete in decimal format from 0 to 1.
	PercentageComplete float32 `protobuf:"fixed32,3,opt,name=percentage_complete,json=percentageComplete,proto3" json:"percentage_complete,omitempty"`
	// Whether or not the project update has failed.
	Failed bool `protobuf:"varint,4,opt,name=failed,proto3" json:"failed,omitempty"`
	// The error message from the failure.
	FailureMessage string `protobuf:"bytes,5,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// Whether or not the project update was canceled.
	Cancelled bool `protobuf:"varint,6,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
}

func (x *ApplyRulesStatusResp) Reset() {
	*x = ApplyRulesStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_response_rules_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRulesStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesStatusResp) ProtoMessage() {}

func (x *ApplyRulesStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_response_rules_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesStatusResp.ProtoReflect.Descriptor instead.
func (*ApplyRulesStatusResp) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_response_rules_proto_rawDescGZIP(), []int{7}
}

func (x *ApplyRulesStatusResp) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ApplyRulesStatusResp) GetEstimatedTimeComplete() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedTimeComplete
	}
	return nil
}

func (x *ApplyRulesStatusResp) GetPercentageComplete() float32 {
	if x != nil {
		return x.PercentageComplete
	}
	return 0
}

func (x *ApplyRulesStatusResp) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

func (x *ApplyRulesStatusResp) GetFailureMessage() string {
	if x != nil {
		return x.FailureMessage
	}
	return ""
}

func (x *ApplyRulesStatusResp) GetCancelled() bool {
	if x != nil {
		return x.Cancelled
	}
	return false
}

var File_external_iam_v2_response_rules_proto protoreflect.FileDescriptor

var file_external_iam_v2_response_rules_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32,
	0x1a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x3a, 0xef, 0x01, 0x92, 0x41, 0xeb,
	0x01, 0x32, 0xe8, 0x01, 0x12, 0xe5, 0x01, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3a, 0x20, 0x22, 0x4e, 0x4f, 0x44, 0x45, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22,
	0x3a, 0x20, 0x22, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x22, 0x2c, 0x20, 0x22,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x69, 0x6f, 0x22, 0x5d, 0x7d, 0x5d, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x3a, 0x20, 0x22, 0x53, 0x54, 0x41, 0x47, 0x45, 0x44, 0x22, 0x7d, 0x22, 0xa5, 0x02, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x32, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x3a, 0xde, 0x01, 0x92, 0x41, 0xda, 0x01, 0x32, 0xd7, 0x01, 0x12, 0xd4, 0x01,
	0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d,
	0x72, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d,
	0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20,
	0x22, 0x4d, 0x79, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x52, 0x75, 0x6c, 0x65,
	0x22, 0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x4f, 0x44, 0x45,
	0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a,
	0x20, 0x5b, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20,
	0x22, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x22, 0x2c, 0x20, 0x22,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x53, 0x22, 0x2c, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b,
	0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x22, 0x5d, 0x7d, 0x5d, 0x2c,
	0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x53, 0x54, 0x41, 0x47,
	0x45, 0x44, 0x22, 0x7d, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x3a, 0xdf, 0x01, 0x92, 0x41, 0xdb, 0x01, 0x32,
	0xd8, 0x01, 0x12, 0xd5, 0x01, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x20, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20,
	0x22, 0x4e, 0x4f, 0x44, 0x45, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3a, 0x20,
	0x22, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x22, 0x2c, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x22, 0x5d, 0x7d, 0x5d, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20,
	0x22, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x22, 0x7d, 0x22, 0x82, 0x05, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x3a, 0xea, 0x03, 0x92, 0x41, 0xe6, 0x03, 0x32, 0xe3, 0x03, 0x12, 0xe0, 0x03, 0x7b,
	0x22, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a,
	0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x2c,
	0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x64, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x4f, 0x44, 0x45, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x43, 0x48, 0x45, 0x46, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x22, 0x2c, 0x20, 0x22,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x22, 0x5d, 0x7d, 0x5d, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x22, 0x7d, 0x2c,
	0x20, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2d, 0x72, 0x75, 0x6c, 0x65, 0x2d, 0x32, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x32, 0x6e, 0x64, 0x20, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3a, 0x20, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x4f, 0x46, 0x22, 0x2c, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20,
	0x5b, 0x22, 0x65, 0x61, 0x73, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x77, 0x65, 0x73, 0x74, 0x22, 0x5d,
	0x7d, 0x5d, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x41,
	0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x22, 0x7d, 0x5d, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x22, 0x7d, 0x22,
	0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x22, 0xbd, 0x03, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x52, 0x0a, 0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x12, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x64, 0x3a, 0xaa, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x32, 0xa3, 0x01, 0x12, 0xa0, 0x01,
	0x7b, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x32, 0x30, 0x32, 0x30, 0x2d, 0x30, 0x33, 0x2d, 0x32, 0x30, 0x54, 0x31, 0x39, 0x3a,
	0x32, 0x34, 0x3a, 0x35, 0x35, 0x5a, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x3a, 0x20,
	0x30, 0x2e, 0x35, 0x2c, 0x20, 0x22, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x2c, 0x20, 0x22, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x7d,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_external_iam_v2_response_rules_proto_rawDescOnce sync.Once
	file_external_iam_v2_response_rules_proto_rawDescData = file_external_iam_v2_response_rules_proto_rawDesc
)

func file_external_iam_v2_response_rules_proto_rawDescGZIP() []byte {
	file_external_iam_v2_response_rules_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_response_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_iam_v2_response_rules_proto_rawDescData)
	})
	return file_external_iam_v2_response_rules_proto_rawDescData
}

var file_external_iam_v2_response_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_external_iam_v2_response_rules_proto_goTypes = []interface{}{
	(*CreateRuleResp)(nil),          // 0: chef.automate.api.iam.v2.CreateRuleResp
	(*UpdateRuleResp)(nil),          // 1: chef.automate.api.iam.v2.UpdateRuleResp
	(*GetRuleResp)(nil),             // 2: chef.automate.api.iam.v2.GetRuleResp
	(*ListRulesForProjectResp)(nil), // 3: chef.automate.api.iam.v2.ListRulesForProjectResp
	(*DeleteRuleResp)(nil),          // 4: chef.automate.api.iam.v2.DeleteRuleResp
	(*ApplyRulesStartResp)(nil),     // 5: chef.automate.api.iam.v2.ApplyRulesStartResp
	(*ApplyRulesCancelResp)(nil),    // 6: chef.automate.api.iam.v2.ApplyRulesCancelResp
	(*ApplyRulesStatusResp)(nil),    // 7: chef.automate.api.iam.v2.ApplyRulesStatusResp
	(*common.Rule)(nil),             // 8: chef.automate.api.iam.v2.Rule
	(common.ProjectRulesStatus)(0),  // 9: chef.automate.api.iam.v2.ProjectRulesStatus
	(*timestamppb.Timestamp)(nil),   // 10: google.protobuf.Timestamp
}
var file_external_iam_v2_response_rules_proto_depIdxs = []int32{
	8,  // 0: chef.automate.api.iam.v2.CreateRuleResp.rule:type_name -> chef.automate.api.iam.v2.Rule
	8,  // 1: chef.automate.api.iam.v2.UpdateRuleResp.rule:type_name -> chef.automate.api.iam.v2.Rule
	8,  // 2: chef.automate.api.iam.v2.GetRuleResp.rule:type_name -> chef.automate.api.iam.v2.Rule
	8,  // 3: chef.automate.api.iam.v2.ListRulesForProjectResp.rules:type_name -> chef.automate.api.iam.v2.Rule
	9,  // 4: chef.automate.api.iam.v2.ListRulesForProjectResp.status:type_name -> chef.automate.api.iam.v2.ProjectRulesStatus
	10, // 5: chef.automate.api.iam.v2.ApplyRulesStatusResp.estimated_time_complete:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_external_iam_v2_response_rules_proto_init() }
func file_external_iam_v2_response_rules_proto_init() {
	if File_external_iam_v2_response_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_iam_v2_response_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRuleResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRulesForProjectResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRuleResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRulesStartResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRulesCancelResp); i {
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
		file_external_iam_v2_response_rules_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRulesStatusResp); i {
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
			RawDescriptor: file_external_iam_v2_response_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_response_rules_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_response_rules_proto_depIdxs,
		MessageInfos:      file_external_iam_v2_response_rules_proto_msgTypes,
	}.Build()
	File_external_iam_v2_response_rules_proto = out.File
	file_external_iam_v2_response_rules_proto_rawDesc = nil
	file_external_iam_v2_response_rules_proto_goTypes = nil
	file_external_iam_v2_response_rules_proto_depIdxs = nil
}
