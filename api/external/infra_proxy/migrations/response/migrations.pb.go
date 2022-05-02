// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: external/infra_proxy/migrations/response/migrations.proto

package response

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

type UploadZipFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Success responce for zip upload.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Migration ID
	MigrationId string `protobuf:"bytes,2,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	// Error data in case of failure.
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *UploadZipFileResponse) Reset() {
	*x = UploadZipFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadZipFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadZipFileResponse) ProtoMessage() {}

func (x *UploadZipFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadZipFileResponse.ProtoReflect.Descriptor instead.
func (*UploadZipFileResponse) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{0}
}

func (x *UploadZipFileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadZipFileResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *UploadZipFileResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetMigrationStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Migration ID.
	MigrationId string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	// Migration type
	MigrationType string `protobuf:"bytes,2,opt,name=migration_type,json=migrationType,proto3" json:"migration_type,omitempty"`
	// Migration status
	MigrationStatus string `protobuf:"bytes,3,opt,name=migration_status,json=migrationStatus,proto3" json:"migration_status,omitempty"`
}

func (x *GetMigrationStatusResponse) Reset() {
	*x = GetMigrationStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMigrationStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMigrationStatusResponse) ProtoMessage() {}

func (x *GetMigrationStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMigrationStatusResponse.ProtoReflect.Descriptor instead.
func (*GetMigrationStatusResponse) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{1}
}

func (x *GetMigrationStatusResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *GetMigrationStatusResponse) GetMigrationType() string {
	if x != nil {
		return x.MigrationType
	}
	return ""
}

func (x *GetMigrationStatusResponse) GetMigrationStatus() string {
	if x != nil {
		return x.MigrationStatus
	}
	return ""
}

type CancelMigrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  migration ID
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Error message
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CancelMigrationResponse) Reset() {
	*x = CancelMigrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMigrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMigrationResponse) ProtoMessage() {}

func (x *CancelMigrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMigrationResponse.ProtoReflect.Descriptor instead.
func (*CancelMigrationResponse) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{2}
}

func (x *CancelMigrationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CancelMigrationResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetStagedDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Migration ID.
	MigrationId string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	// Staged data
	StagedData *StagedData `protobuf:"bytes,2,opt,name=staged_data,json=stagedData,proto3" json:"staged_data,omitempty"`
}

func (x *GetStagedDataResponse) Reset() {
	*x = GetStagedDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStagedDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStagedDataResponse) ProtoMessage() {}

func (x *GetStagedDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStagedDataResponse.ProtoReflect.Descriptor instead.
func (*GetStagedDataResponse) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{3}
}

func (x *GetStagedDataResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *GetStagedDataResponse) GetStagedData() *StagedData {
	if x != nil {
		return x.StagedData
	}
	return nil
}

type StagedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of orgs to migrate
	OrgsToMigrate int32 `protobuf:"varint,1,opt,name=orgs_to_migrate,json=orgsToMigrate,proto3" json:"orgs_to_migrate,omitempty"`
	// Number of orgs to skip
	OrgsToSkip int32 `protobuf:"varint,2,opt,name=orgs_to_skip,json=orgsToSkip,proto3" json:"orgs_to_skip,omitempty"`
	// Number of orgs to update
	OrgsToUpdate int32 `protobuf:"varint,3,opt,name=orgs_to_update,json=orgsToUpdate,proto3" json:"orgs_to_update,omitempty"`
	// Number of orgs to delete
	OrgsToDelete int32 `protobuf:"varint,4,opt,name=orgs_to_delete,json=orgsToDelete,proto3" json:"orgs_to_delete,omitempty"`
	// Users
	Users []*User `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *StagedData) Reset() {
	*x = StagedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StagedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StagedData) ProtoMessage() {}

func (x *StagedData) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StagedData.ProtoReflect.Descriptor instead.
func (*StagedData) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{4}
}

func (x *StagedData) GetOrgsToMigrate() int32 {
	if x != nil {
		return x.OrgsToMigrate
	}
	return 0
}

func (x *StagedData) GetOrgsToSkip() int32 {
	if x != nil {
		return x.OrgsToSkip
	}
	return 0
}

func (x *StagedData) GetOrgsToUpdate() int32 {
	if x != nil {
		return x.OrgsToUpdate
	}
	return 0
}

func (x *StagedData) GetOrgsToDelete() int32 {
	if x != nil {
		return x.OrgsToDelete
	}
	return 0
}

func (x *StagedData) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's username
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// User's email ID
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// User's display name
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User's first name
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// User's last name
	LastName string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// User's middle name
	MiddleName string `protobuf:"bytes,6,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	// User's username in automate
	AutomateUsername string `protobuf:"bytes,7,opt,name=automate_username,json=automateUsername,proto3" json:"automate_username,omitempty"`
	// Local or ldap user
	Connector string `protobuf:"bytes,8,opt,name=connector,proto3" json:"connector,omitempty"`
	// IsConflicting for user's existance in db
	IsConflicting bool `protobuf:"varint,9,opt,name=is_conflicting,json=isConflicting,proto3" json:"is_conflicting,omitempty"`
	// user is admin or not
	IsAdmin bool `protobuf:"varint,10,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	//Local user hashed password
	HashPassword string `protobuf:"bytes,11,opt,name=hash_password,json=hashPassword,proto3" json:"hash_password,omitempty"`
	// Local User actionops
	ActionOps int32 `protobuf:"varint,12,opt,name=action_ops,json=actionOps,proto3" json:"action_ops,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetAutomateUsername() string {
	if x != nil {
		return x.AutomateUsername
	}
	return ""
}

func (x *User) GetConnector() string {
	if x != nil {
		return x.Connector
	}
	return ""
}

func (x *User) GetIsConflicting() bool {
	if x != nil {
		return x.IsConflicting
	}
	return false
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *User) GetHashPassword() string {
	if x != nil {
		return x.HashPassword
	}
	return ""
}

func (x *User) GetActionOps() int32 {
	if x != nil {
		return x.ActionOps
	}
	return 0
}

type ConfirmPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Migration ID
	MigrationId string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
}

func (x *ConfirmPreview) Reset() {
	*x = ConfirmPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPreview) ProtoMessage() {}

func (x *ConfirmPreview) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPreview.ProtoReflect.Descriptor instead.
func (*ConfirmPreview) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP(), []int{6}
}

func (x *ConfirmPreview) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

var File_external_infra_proxy_migrations_response_migrations_proto protoreflect.FileDescriptor

var file_external_infra_proxy_migrations_response_migrations_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6c,
	0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x69, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x91, 0x01, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x4b, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x9a, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0xf1, 0x01, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x67,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x73, 0x54, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x6b, 0x69,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x73, 0x54, 0x6f, 0x53,
	0x6b, 0x69, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x73, 0x54, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x72, 0x67,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x73, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x4d, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x89,
	0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c,
	0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68,
	0x61, 0x73, 0x68, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x73, 0x22, 0x33, 0x0a, 0x0e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_migrations_response_migrations_proto_rawDescOnce sync.Once
	file_external_infra_proxy_migrations_response_migrations_proto_rawDescData = file_external_infra_proxy_migrations_response_migrations_proto_rawDesc
)

func file_external_infra_proxy_migrations_response_migrations_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_migrations_response_migrations_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_migrations_response_migrations_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_migrations_response_migrations_proto_rawDescData)
	})
	return file_external_infra_proxy_migrations_response_migrations_proto_rawDescData
}

var file_external_infra_proxy_migrations_response_migrations_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_external_infra_proxy_migrations_response_migrations_proto_goTypes = []interface{}{
	(*UploadZipFileResponse)(nil),      // 0: chef.automate.api.infra_proxy.migrations.response.UploadZipFileResponse
	(*GetMigrationStatusResponse)(nil), // 1: chef.automate.api.infra_proxy.migrations.response.GetMigrationStatusResponse
	(*CancelMigrationResponse)(nil),    // 2: chef.automate.api.infra_proxy.migrations.response.CancelMigrationResponse
	(*GetStagedDataResponse)(nil),      // 3: chef.automate.api.infra_proxy.migrations.response.GetStagedDataResponse
	(*StagedData)(nil),                 // 4: chef.automate.api.infra_proxy.migrations.response.StagedData
	(*User)(nil),                       // 5: chef.automate.api.infra_proxy.migrations.response.User
	(*ConfirmPreview)(nil),             // 6: chef.automate.api.infra_proxy.migrations.response.ConfirmPreview
}
var file_external_infra_proxy_migrations_response_migrations_proto_depIdxs = []int32{
	4, // 0: chef.automate.api.infra_proxy.migrations.response.GetStagedDataResponse.staged_data:type_name -> chef.automate.api.infra_proxy.migrations.response.StagedData
	5, // 1: chef.automate.api.infra_proxy.migrations.response.StagedData.users:type_name -> chef.automate.api.infra_proxy.migrations.response.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_migrations_response_migrations_proto_init() }
func file_external_infra_proxy_migrations_response_migrations_proto_init() {
	if File_external_infra_proxy_migrations_response_migrations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadZipFileResponse); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMigrationStatusResponse); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMigrationResponse); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStagedDataResponse); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StagedData); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_external_infra_proxy_migrations_response_migrations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPreview); i {
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
			RawDescriptor: file_external_infra_proxy_migrations_response_migrations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_migrations_response_migrations_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_migrations_response_migrations_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_migrations_response_migrations_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_migrations_response_migrations_proto = out.File
	file_external_infra_proxy_migrations_response_migrations_proto_rawDesc = nil
	file_external_infra_proxy_migrations_response_migrations_proto_goTypes = nil
	file_external_infra_proxy_migrations_response_migrations_proto_depIdxs = nil
}
