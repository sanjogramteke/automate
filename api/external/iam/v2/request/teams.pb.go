// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/iam/v2/request/teams.proto

package request

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type ListTeamsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTeamsReq) Reset() {
	*x = ListTeamsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsReq) ProtoMessage() {}

func (x *ListTeamsReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsReq.ProtoReflect.Descriptor instead.
func (*ListTeamsReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{0}
}

type GetTeamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTeamReq) Reset() {
	*x = GetTeamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamReq) ProtoMessage() {}

func (x *GetTeamReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamReq.ProtoReflect.Descriptor instead.
func (*GetTeamReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{1}
}

func (x *GetTeamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTeamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *CreateTeamReq) Reset() {
	*x = CreateTeamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamReq) ProtoMessage() {}

func (x *CreateTeamReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamReq.ProtoReflect.Descriptor instead.
func (*CreateTeamReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTeamReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTeamReq) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type UpdateTeamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *UpdateTeamReq) Reset() {
	*x = UpdateTeamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamReq) ProtoMessage() {}

func (x *UpdateTeamReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamReq.ProtoReflect.Descriptor instead.
func (*UpdateTeamReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTeamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeamReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTeamReq) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type DeleteTeamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTeamReq) Reset() {
	*x = DeleteTeamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamReq) ProtoMessage() {}

func (x *DeleteTeamReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamReq.ProtoReflect.Descriptor instead.
func (*DeleteTeamReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTeamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddTeamMembersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MembershipIds []string `protobuf:"bytes,2,rep,name=membership_ids,json=membershipIds,proto3" json:"membership_ids,omitempty"`
}

func (x *AddTeamMembersReq) Reset() {
	*x = AddTeamMembersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTeamMembersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTeamMembersReq) ProtoMessage() {}

func (x *AddTeamMembersReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTeamMembersReq.ProtoReflect.Descriptor instead.
func (*AddTeamMembersReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{5}
}

func (x *AddTeamMembersReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddTeamMembersReq) GetMembershipIds() []string {
	if x != nil {
		return x.MembershipIds
	}
	return nil
}

type GetTeamMembershipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTeamMembershipReq) Reset() {
	*x = GetTeamMembershipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamMembershipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamMembershipReq) ProtoMessage() {}

func (x *GetTeamMembershipReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamMembershipReq.ProtoReflect.Descriptor instead.
func (*GetTeamMembershipReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{6}
}

func (x *GetTeamMembershipReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveTeamMembersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MembershipIds []string `protobuf:"bytes,2,rep,name=membership_ids,json=membershipIds,proto3" json:"membership_ids,omitempty"`
}

func (x *RemoveTeamMembersReq) Reset() {
	*x = RemoveTeamMembersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTeamMembersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTeamMembersReq) ProtoMessage() {}

func (x *RemoveTeamMembersReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTeamMembersReq.ProtoReflect.Descriptor instead.
func (*RemoveTeamMembersReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveTeamMembersReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveTeamMembersReq) GetMembershipIds() []string {
	if x != nil {
		return x.MembershipIds
	}
	return nil
}

type GetTeamsForMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MembershipId string `protobuf:"bytes,1,opt,name=membership_id,json=membershipId,proto3" json:"membership_id,omitempty"`
}

func (x *GetTeamsForMemberReq) Reset() {
	*x = GetTeamsForMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamsForMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamsForMemberReq) ProtoMessage() {}

func (x *GetTeamsForMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamsForMemberReq.ProtoReflect.Descriptor instead.
func (*GetTeamsForMemberReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{8}
}

func (x *GetTeamsForMemberReq) GetMembershipId() string {
	if x != nil {
		return x.MembershipId
	}
	return ""
}

type ApplyV2DataMigrationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyV2DataMigrationsReq) Reset() {
	*x = ApplyV2DataMigrationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyV2DataMigrationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyV2DataMigrationsReq) ProtoMessage() {}

func (x *ApplyV2DataMigrationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyV2DataMigrationsReq.ProtoReflect.Descriptor instead.
func (*ApplyV2DataMigrationsReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{9}
}

type ResetAllTeamProjectsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetAllTeamProjectsReq) Reset() {
	*x = ResetAllTeamProjectsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_request_teams_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetAllTeamProjectsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetAllTeamProjectsReq) ProtoMessage() {}

func (x *ResetAllTeamProjectsReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_teams_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetAllTeamProjectsReq.ProtoReflect.Descriptor instead.
func (*ResetAllTeamProjectsReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_teams_proto_rawDescGZIP(), []int{10}
}

var File_external_iam_v2_request_teams_proto protoreflect.FileDescriptor

var file_external_iam_v2_request_teams_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x28, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x0a, 0x92, 0x41, 0x07,
	0x0a, 0x05, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x3a, 0x67, 0x92, 0x41, 0x64, 0x0a, 0x0c,
	0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x54, 0x12, 0x52,
	0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x6d, 0x79, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2d,
	0x69, 0x64, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79,
	0x20, 0x54, 0x65, 0x73, 0x74, 0x20, 0x54, 0x65, 0x61, 0x6d, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x31, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x22,
	0x5d, 0x7d, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x3a, 0x67, 0x92, 0x41, 0x64, 0x0a, 0x0c, 0xd2, 0x01, 0x02, 0x69, 0x64,
	0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x54, 0x12, 0x52, 0x7b, 0x22, 0x69, 0x64, 0x22,
	0x3a, 0x20, 0x22, 0x6d, 0x79, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x69, 0x64, 0x22, 0x2c, 0x20,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x54, 0x65, 0x73, 0x74,
	0x20, 0x54, 0x65, 0x61, 0x6d, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x31, 0x22, 0x2c,
	0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x22, 0x5d, 0x7d, 0x22, 0x42, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x21,
	0x92, 0x41, 0x1e, 0x0a, 0x05, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x32, 0x15, 0x12, 0x13, 0x7b, 0x22,
	0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22,
	0x7d, 0x22, 0xe1, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x73, 0x3a, 0x94,
	0x01, 0x92, 0x41, 0x90, 0x01, 0x0a, 0x16, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x32, 0x76, 0x12,
	0x74, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x22,
	0x2c, 0x20, 0x22, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64,
	0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x35, 0x32, 0x37, 0x65, 0x64, 0x39, 0x36, 0x66, 0x2d, 0x32,
	0x65, 0x63, 0x62, 0x2d, 0x34, 0x66, 0x38, 0x66, 0x2d, 0x61, 0x62, 0x64, 0x37, 0x2d, 0x30, 0x62,
	0x66, 0x36, 0x35, 0x31, 0x31, 0x34, 0x35, 0x39, 0x61, 0x63, 0x22, 0x2c, 0x20, 0x22, 0x33, 0x35,
	0x33, 0x61, 0x36, 0x32, 0x64, 0x34, 0x2d, 0x38, 0x35, 0x66, 0x61, 0x2d, 0x34, 0x34, 0x32, 0x33,
	0x2d, 0x62, 0x31, 0x32, 0x61, 0x2d, 0x66, 0x36, 0x36, 0x30, 0x38, 0x61, 0x35, 0x36, 0x32, 0x61,
	0x65, 0x39, 0x22, 0x5d, 0x7d, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x0a, 0x92,
	0x41, 0x07, 0x0a, 0x05, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x14, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x73, 0x3a, 0x94, 0x01, 0x92, 0x41, 0x90, 0x01,
	0x0a, 0x16, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x32, 0x76, 0x12, 0x74, 0x7b, 0x22, 0x69, 0x64,
	0x22, 0x3a, 0x20, 0x22, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x22, 0x2c, 0x20, 0x22, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x22, 0x3a, 0x20, 0x5b,
	0x22, 0x35, 0x32, 0x37, 0x65, 0x64, 0x39, 0x36, 0x66, 0x2d, 0x32, 0x65, 0x63, 0x62, 0x2d, 0x34,
	0x66, 0x38, 0x66, 0x2d, 0x61, 0x62, 0x64, 0x37, 0x2d, 0x30, 0x62, 0x66, 0x36, 0x35, 0x31, 0x31,
	0x34, 0x35, 0x39, 0x61, 0x63, 0x22, 0x2c, 0x20, 0x22, 0x33, 0x35, 0x33, 0x61, 0x36, 0x32, 0x64,
	0x34, 0x2d, 0x38, 0x35, 0x66, 0x61, 0x2d, 0x34, 0x34, 0x32, 0x33, 0x2d, 0x62, 0x31, 0x32, 0x61,
	0x2d, 0x66, 0x36, 0x36, 0x30, 0x38, 0x61, 0x35, 0x36, 0x32, 0x61, 0x65, 0x39, 0x22, 0x5d, 0x7d,
	0x22, 0x8f, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x46, 0x6f, 0x72,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x3a, 0x52,
	0x92, 0x41, 0x4f, 0x0a, 0x10, 0xd2, 0x01, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x69, 0x64, 0x32, 0x3b, 0x12, 0x39, 0x7b, 0x22, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x35, 0x32, 0x37, 0x65,
	0x64, 0x39, 0x36, 0x66, 0x2d, 0x32, 0x65, 0x63, 0x62, 0x2d, 0x34, 0x66, 0x38, 0x66, 0x2d, 0x61,
	0x62, 0x64, 0x37, 0x2d, 0x30, 0x62, 0x66, 0x36, 0x35, 0x31, 0x31, 0x34, 0x35, 0x39, 0x61, 0x63,
	0x22, 0x7d, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x56, 0x32, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x22, 0x19,
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_iam_v2_request_teams_proto_rawDescOnce sync.Once
	file_external_iam_v2_request_teams_proto_rawDescData = file_external_iam_v2_request_teams_proto_rawDesc
)

func file_external_iam_v2_request_teams_proto_rawDescGZIP() []byte {
	file_external_iam_v2_request_teams_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_request_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_iam_v2_request_teams_proto_rawDescData)
	})
	return file_external_iam_v2_request_teams_proto_rawDescData
}

var file_external_iam_v2_request_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_external_iam_v2_request_teams_proto_goTypes = []interface{}{
	(*ListTeamsReq)(nil),             // 0: chef.automate.api.iam.v2.ListTeamsReq
	(*GetTeamReq)(nil),               // 1: chef.automate.api.iam.v2.GetTeamReq
	(*CreateTeamReq)(nil),            // 2: chef.automate.api.iam.v2.CreateTeamReq
	(*UpdateTeamReq)(nil),            // 3: chef.automate.api.iam.v2.UpdateTeamReq
	(*DeleteTeamReq)(nil),            // 4: chef.automate.api.iam.v2.DeleteTeamReq
	(*AddTeamMembersReq)(nil),        // 5: chef.automate.api.iam.v2.AddTeamMembersReq
	(*GetTeamMembershipReq)(nil),     // 6: chef.automate.api.iam.v2.GetTeamMembershipReq
	(*RemoveTeamMembersReq)(nil),     // 7: chef.automate.api.iam.v2.RemoveTeamMembersReq
	(*GetTeamsForMemberReq)(nil),     // 8: chef.automate.api.iam.v2.GetTeamsForMemberReq
	(*ApplyV2DataMigrationsReq)(nil), // 9: chef.automate.api.iam.v2.ApplyV2DataMigrationsReq
	(*ResetAllTeamProjectsReq)(nil),  // 10: chef.automate.api.iam.v2.ResetAllTeamProjectsReq
}
var file_external_iam_v2_request_teams_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_iam_v2_request_teams_proto_init() }
func file_external_iam_v2_request_teams_proto_init() {
	if File_external_iam_v2_request_teams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_iam_v2_request_teams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeamReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTeamMembersReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamMembershipReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTeamMembersReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamsForMemberReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyV2DataMigrationsReq); i {
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
		file_external_iam_v2_request_teams_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetAllTeamProjectsReq); i {
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
			RawDescriptor: file_external_iam_v2_request_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_request_teams_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_request_teams_proto_depIdxs,
		MessageInfos:      file_external_iam_v2_request_teams_proto_msgTypes,
	}.Build()
	File_external_iam_v2_request_teams_proto = out.File
	file_external_iam_v2_request_teams_proto_rawDesc = nil
	file_external_iam_v2_request_teams_proto_goTypes = nil
	file_external_iam_v2_request_teams_proto_depIdxs = nil
}
