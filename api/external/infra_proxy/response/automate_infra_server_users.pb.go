// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/infra_proxy/response/automate_infra_server_users.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AutomateInfraServerUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Automate infra server users list
	Users []*AutomateInfraServerUsersListItem `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *AutomateInfraServerUsers) Reset() {
	*x = AutomateInfraServerUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraServerUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraServerUsers) ProtoMessage() {}

func (x *AutomateInfraServerUsers) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraServerUsers.ProtoReflect.Descriptor instead.
func (*AutomateInfraServerUsers) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescGZIP(), []int{0}
}

func (x *AutomateInfraServerUsers) GetUsers() []*AutomateInfraServerUsersListItem {
	if x != nil {
		return x.Users
	}
	return nil
}

type AutomateInfraServerOrgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization list.
	Orgs []*AutomateInfraServerOrg `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty"`
}

func (x *AutomateInfraServerOrgs) Reset() {
	*x = AutomateInfraServerOrgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraServerOrgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraServerOrgs) ProtoMessage() {}

func (x *AutomateInfraServerOrgs) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraServerOrgs.ProtoReflect.Descriptor instead.
func (*AutomateInfraServerOrgs) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescGZIP(), []int{1}
}

func (x *AutomateInfraServerOrgs) GetOrgs() []*AutomateInfraServerOrg {
	if x != nil {
		return x.Orgs
	}
	return nil
}

type AutomateInfraServerUsersListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User server id
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// User infra server username
	InfraServerUsername string `protobuf:"bytes,3,opt,name=infra_server_username,json=infraServerUsername,proto3" json:"infra_server_username,omitempty"`
	// User credential id
	CredentialId string `protobuf:"bytes,4,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty"`
	// User connector
	Connector string `protobuf:"bytes,5,opt,name=connector,proto3" json:"connector,omitempty"`
	// User automate user id
	AutomateUserId string `protobuf:"bytes,6,opt,name=automate_user_id,json=automateUserId,proto3" json:"automate_user_id,omitempty"`
	// User is server admin or not
	IsServerAdmin bool `protobuf:"varint,7,opt,name=is_server_admin,json=isServerAdmin,proto3" json:"is_server_admin,omitempty"`
}

func (x *AutomateInfraServerUsersListItem) Reset() {
	*x = AutomateInfraServerUsersListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraServerUsersListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraServerUsersListItem) ProtoMessage() {}

func (x *AutomateInfraServerUsersListItem) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraServerUsersListItem.ProtoReflect.Descriptor instead.
func (*AutomateInfraServerUsersListItem) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescGZIP(), []int{2}
}

func (x *AutomateInfraServerUsersListItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetInfraServerUsername() string {
	if x != nil {
		return x.InfraServerUsername
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetCredentialId() string {
	if x != nil {
		return x.CredentialId
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetConnector() string {
	if x != nil {
		return x.Connector
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetAutomateUserId() string {
	if x != nil {
		return x.AutomateUserId
	}
	return ""
}

func (x *AutomateInfraServerUsersListItem) GetIsServerAdmin() bool {
	if x != nil {
		return x.IsServerAdmin
	}
	return false
}

type AutomateInfraServerOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty"`
	// Chef organization credential ID.
	CredentialId string `protobuf:"bytes,4,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *AutomateInfraServerOrg) Reset() {
	*x = AutomateInfraServerOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraServerOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraServerOrg) ProtoMessage() {}

func (x *AutomateInfraServerOrg) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraServerOrg.ProtoReflect.Descriptor instead.
func (*AutomateInfraServerOrg) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescGZIP(), []int{3}
}

func (x *AutomateInfraServerOrg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AutomateInfraServerOrg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AutomateInfraServerOrg) GetAdminUser() string {
	if x != nil {
		return x.AdminUser
	}
	return ""
}

func (x *AutomateInfraServerOrg) GetCredentialId() string {
	if x != nil {
		return x.CredentialId
	}
	return ""
}

func (x *AutomateInfraServerOrg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AutomateInfraServerOrg) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

var File_external_infra_proxy_response_automate_infra_server_users_proto protoreflect.FileDescriptor

var file_external_infra_proxy_response_automate_infra_server_users_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x26, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x18, 0x41, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x5e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x6d, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x73,
	0x12, 0x52, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x52, 0x04,
	0x6f, 0x72, 0x67, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x20, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0xb9, 0x01, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescOnce sync.Once
	file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescData = file_external_infra_proxy_response_automate_infra_server_users_proto_rawDesc
)

func file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescData)
	})
	return file_external_infra_proxy_response_automate_infra_server_users_proto_rawDescData
}

var file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_external_infra_proxy_response_automate_infra_server_users_proto_goTypes = []interface{}{
	(*AutomateInfraServerUsers)(nil),         // 0: chef.automate.api.infra_proxy.response.AutomateInfraServerUsers
	(*AutomateInfraServerOrgs)(nil),          // 1: chef.automate.api.infra_proxy.response.AutomateInfraServerOrgs
	(*AutomateInfraServerUsersListItem)(nil), // 2: chef.automate.api.infra_proxy.response.AutomateInfraServerUsersListItem
	(*AutomateInfraServerOrg)(nil),           // 3: chef.automate.api.infra_proxy.response.AutomateInfraServerOrg
}
var file_external_infra_proxy_response_automate_infra_server_users_proto_depIdxs = []int32{
	2, // 0: chef.automate.api.infra_proxy.response.AutomateInfraServerUsers.users:type_name -> chef.automate.api.infra_proxy.response.AutomateInfraServerUsersListItem
	3, // 1: chef.automate.api.infra_proxy.response.AutomateInfraServerOrgs.orgs:type_name -> chef.automate.api.infra_proxy.response.AutomateInfraServerOrg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_response_automate_infra_server_users_proto_init() }
func file_external_infra_proxy_response_automate_infra_server_users_proto_init() {
	if File_external_infra_proxy_response_automate_infra_server_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraServerUsers); i {
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
		file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraServerOrgs); i {
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
		file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraServerUsersListItem); i {
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
		file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraServerOrg); i {
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
			RawDescriptor: file_external_infra_proxy_response_automate_infra_server_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_response_automate_infra_server_users_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_response_automate_infra_server_users_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_response_automate_infra_server_users_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_response_automate_infra_server_users_proto = out.File
	file_external_infra_proxy_response_automate_infra_server_users_proto_rawDesc = nil
	file_external_infra_proxy_response_automate_infra_server_users_proto_goTypes = nil
	file_external_infra_proxy_response_automate_infra_server_users_proto_depIdxs = nil
}
