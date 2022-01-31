// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: interservice/infra_proxy/response/servers.proto

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

type CreateServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server.
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" toml:"server,omitempty" mapstructure:"server,omitempty"`
}

func (x *CreateServer) Reset() {
	*x = CreateServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServer) ProtoMessage() {}

func (x *CreateServer) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServer.ProtoReflect.Descriptor instead.
func (*CreateServer) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServer) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type DeleteServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server.
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" toml:"server,omitempty" mapstructure:"server,omitempty"`
}

func (x *DeleteServer) Reset() {
	*x = DeleteServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServer) ProtoMessage() {}

func (x *DeleteServer) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServer.ProtoReflect.Descriptor instead.
func (*DeleteServer) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteServer) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type UpdateServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server.
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" toml:"server,omitempty" mapstructure:"server,omitempty"`
}

func (x *UpdateServer) Reset() {
	*x = UpdateServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServer) ProtoMessage() {}

func (x *UpdateServer) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServer.ProtoReflect.Descriptor instead.
func (*UpdateServer) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateServer) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type GetServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Chef Infra Servers.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty" toml:"servers,omitempty" mapstructure:"servers,omitempty"`
}

func (x *GetServers) Reset() {
	*x = GetServers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServers) ProtoMessage() {}

func (x *GetServers) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServers.ProtoReflect.Descriptor instead.
func (*GetServers) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{3}
}

func (x *GetServers) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type GetServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server.
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" toml:"server,omitempty" mapstructure:"server,omitempty"`
}

func (x *GetServer) Reset() {
	*x = GetServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServer) ProtoMessage() {}

func (x *GetServer) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServer.ProtoReflect.Descriptor instead.
func (*GetServer) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{4}
}

func (x *GetServer) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef Infra Server FQDN.
	Fqdn string `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	// Chef Infra Server IP address.
	IpAddress string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty" toml:"ip_address,omitempty" mapstructure:"ip_address,omitempty"`
	// Chef organizations count associated with Chef Infra Server.
	OrgsCount int32 `protobuf:"varint,6,opt,name=orgs_count,json=orgsCount,proto3" json:"orgs_count,omitempty" toml:"orgs_count,omitempty" mapstructure:"orgs_count,omitempty"`
	//Migration Id for active Chef Infra server migration
	MigrationId string `protobuf:"bytes,7,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty" toml:"migration_id,omitempty" mapstructure:"migration_id,omitempty"`
	//Migration Phase for active Chef Infra server migration
	MigrationType string `protobuf:"bytes,8,opt,name=migration_type,json=migrationType,proto3" json:"migration_type,omitempty" toml:"migration_type,omitempty" mapstructure:"migration_type,omitempty"`
	//Migration type for active Chef Infra server migration
	MigrationStatus string `protobuf:"bytes,9,opt,name=migration_status,json=migrationStatus,proto3" json:"migration_status,omitempty" toml:"migration_status,omitempty" mapstructure:"migration_status,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{5}
}

func (x *Server) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *Server) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *Server) GetOrgsCount() int32 {
	if x != nil {
		return x.OrgsCount
	}
	return 0
}

func (x *Server) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *Server) GetMigrationType() string {
	if x != nil {
		return x.MigrationType
	}
	return ""
}

func (x *Server) GetMigrationStatus() string {
	if x != nil {
		return x.MigrationStatus
	}
	return ""
}

type GetServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Server Status
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
	// Chef Server Upstream
	Upstreams map[string]string `protobuf:"bytes,2,rep,name=upstreams,proto3" json:"upstreams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" toml:"upstreams,omitempty" mapstructure:"upstreams,omitempty"`
	// Chef Server Keygem
	Keygen map[string]int32 `protobuf:"bytes,3,rep,name=keygen,proto3" json:"keygen,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" toml:"keygen,omitempty" mapstructure:"keygen,omitempty"`
}

func (x *GetServerStatus) Reset() {
	*x = GetServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerStatus) ProtoMessage() {}

func (x *GetServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerStatus.ProtoReflect.Descriptor instead.
func (*GetServerStatus) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{6}
}

func (x *GetServerStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetServerStatus) GetUpstreams() map[string]string {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

func (x *GetServerStatus) GetKeygen() map[string]int32 {
	if x != nil {
		return x.Keygen
	}
	return nil
}

type ValidateWebuiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Webui key is valid or not
	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty" toml:"valid,omitempty" mapstructure:"valid,omitempty"`
	// webui key validation errors
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" toml:"error,omitempty" mapstructure:"error,omitempty"`
}

func (x *ValidateWebuiKey) Reset() {
	*x = ValidateWebuiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWebuiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWebuiKey) ProtoMessage() {}

func (x *ValidateWebuiKey) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateWebuiKey.ProtoReflect.Descriptor instead.
func (*ValidateWebuiKey) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateWebuiKey) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateWebuiKey) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateWebuiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateWebuiKey) Reset() {
	*x = UpdateWebuiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWebuiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebuiKey) ProtoMessage() {}

func (x *UpdateWebuiKey) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_servers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebuiKey.ProtoReflect.Descriptor instead.
func (*UpdateWebuiKey) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_servers_proto_rawDescGZIP(), []int{8}
}

var File_interservice_infra_proxy_response_servers_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_response_servers_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0x59, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x59, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0xf3, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x71, 0x64, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x72, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xeb, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x67, 0x0a, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x5e, 0x0a, 0x06, 0x6b,
	0x65, 0x79, 0x67, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x1a, 0x3c, 0x0a, 0x0e, 0x55,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4b, 0x65, 0x79,
	0x67, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x57, 0x65, 0x62, 0x75, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65,
	0x62, 0x75, 0x69, 0x4b, 0x65, 0x79, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_response_servers_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_response_servers_proto_rawDescData = file_interservice_infra_proxy_response_servers_proto_rawDesc
)

func file_interservice_infra_proxy_response_servers_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_response_servers_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_response_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_response_servers_proto_rawDescData)
	})
	return file_interservice_infra_proxy_response_servers_proto_rawDescData
}

var file_interservice_infra_proxy_response_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_interservice_infra_proxy_response_servers_proto_goTypes = []interface{}{
	(*CreateServer)(nil),     // 0: chef.automate.domain.infra_proxy.response.CreateServer
	(*DeleteServer)(nil),     // 1: chef.automate.domain.infra_proxy.response.DeleteServer
	(*UpdateServer)(nil),     // 2: chef.automate.domain.infra_proxy.response.UpdateServer
	(*GetServers)(nil),       // 3: chef.automate.domain.infra_proxy.response.GetServers
	(*GetServer)(nil),        // 4: chef.automate.domain.infra_proxy.response.GetServer
	(*Server)(nil),           // 5: chef.automate.domain.infra_proxy.response.Server
	(*GetServerStatus)(nil),  // 6: chef.automate.domain.infra_proxy.response.GetServerStatus
	(*ValidateWebuiKey)(nil), // 7: chef.automate.domain.infra_proxy.response.ValidateWebuiKey
	(*UpdateWebuiKey)(nil),   // 8: chef.automate.domain.infra_proxy.response.UpdateWebuiKey
	nil,                      // 9: chef.automate.domain.infra_proxy.response.GetServerStatus.UpstreamsEntry
	nil,                      // 10: chef.automate.domain.infra_proxy.response.GetServerStatus.KeygenEntry
}
var file_interservice_infra_proxy_response_servers_proto_depIdxs = []int32{
	5,  // 0: chef.automate.domain.infra_proxy.response.CreateServer.server:type_name -> chef.automate.domain.infra_proxy.response.Server
	5,  // 1: chef.automate.domain.infra_proxy.response.DeleteServer.server:type_name -> chef.automate.domain.infra_proxy.response.Server
	5,  // 2: chef.automate.domain.infra_proxy.response.UpdateServer.server:type_name -> chef.automate.domain.infra_proxy.response.Server
	5,  // 3: chef.automate.domain.infra_proxy.response.GetServers.servers:type_name -> chef.automate.domain.infra_proxy.response.Server
	5,  // 4: chef.automate.domain.infra_proxy.response.GetServer.server:type_name -> chef.automate.domain.infra_proxy.response.Server
	9,  // 5: chef.automate.domain.infra_proxy.response.GetServerStatus.upstreams:type_name -> chef.automate.domain.infra_proxy.response.GetServerStatus.UpstreamsEntry
	10, // 6: chef.automate.domain.infra_proxy.response.GetServerStatus.keygen:type_name -> chef.automate.domain.infra_proxy.response.GetServerStatus.KeygenEntry
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_response_servers_proto_init() }
func file_interservice_infra_proxy_response_servers_proto_init() {
	if File_interservice_infra_proxy_response_servers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_response_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServer); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServer); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServer); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServers); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServer); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerStatus); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateWebuiKey); i {
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
		file_interservice_infra_proxy_response_servers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWebuiKey); i {
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
			RawDescriptor: file_interservice_infra_proxy_response_servers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_response_servers_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_response_servers_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_response_servers_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_response_servers_proto = out.File
	file_interservice_infra_proxy_response_servers_proto_rawDesc = nil
	file_interservice_infra_proxy_response_servers_proto_goTypes = nil
	file_interservice_infra_proxy_response_servers_proto_depIdxs = nil
}
