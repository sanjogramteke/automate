// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/infra_proxy/request/servers.proto

package request

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

type CreateServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef Infra Server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Chef Infra Server FQDN.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Chef Infra Server IP address.
	IpAddress string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Chef Infra Server Webui key.
	WebuiKey string `protobuf:"bytes,5,opt,name=webui_key,json=webuiKey,proto3" json:"webui_key,omitempty"`
}

func (x *CreateServer) Reset() {
	*x = CreateServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServer) ProtoMessage() {}

func (x *CreateServer) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[0]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServer) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *CreateServer) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *CreateServer) GetWebuiKey() string {
	if x != nil {
		return x.WebuiKey
	}
	return ""
}

type UpdateServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef Infra Server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Chef Infra Server FQDN.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Chef Infra Server IP address.
	IpAddress string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *UpdateServer) Reset() {
	*x = UpdateServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServer) ProtoMessage() {}

func (x *UpdateServer) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[1]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateServer) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *UpdateServer) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type DeleteServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteServer) Reset() {
	*x = DeleteServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServer) ProtoMessage() {}

func (x *DeleteServer) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[2]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServers) Reset() {
	*x = GetServers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServers) ProtoMessage() {}

func (x *GetServers) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[3]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{3}
}

type GetServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetServer) Reset() {
	*x = GetServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServer) ProtoMessage() {}

func (x *GetServer) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[4]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{4}
}

func (x *GetServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server FQDN
	Fqdn string `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
}

func (x *GetServerStatus) Reset() {
	*x = GetServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerStatus) ProtoMessage() {}

func (x *GetServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[5]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{5}
}

func (x *GetServerStatus) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

type ValidateWebuiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef Server FQDN
	Fqdn string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Optional Chef Server Webui Key
	WebuiKey string `protobuf:"bytes,3,opt,name=webui_key,json=webuiKey,proto3" json:"webui_key,omitempty"`
}

func (x *ValidateWebuiKey) Reset() {
	*x = ValidateWebuiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_servers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWebuiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWebuiKey) ProtoMessage() {}

func (x *ValidateWebuiKey) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_servers_proto_msgTypes[6]
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
	return file_external_infra_proxy_request_servers_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateWebuiKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ValidateWebuiKey) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *ValidateWebuiKey) GetWebuiKey() string {
	if x != nil {
		return x.WebuiKey
	}
	return ""
}

var File_external_infra_proxy_request_servers_proto protoreflect.FileDescriptor

var file_external_infra_proxy_request_servers_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77,
	0x65, 0x62, 0x75, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x65, 0x62, 0x75, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x65, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x71, 0x64, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x1b, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64,
	0x6e, 0x22, 0x53, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62,
	0x75, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x62,
	0x75, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65,
	0x62, 0x75, 0x69, 0x4b, 0x65, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_request_servers_proto_rawDescOnce sync.Once
	file_external_infra_proxy_request_servers_proto_rawDescData = file_external_infra_proxy_request_servers_proto_rawDesc
)

func file_external_infra_proxy_request_servers_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_request_servers_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_request_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_request_servers_proto_rawDescData)
	})
	return file_external_infra_proxy_request_servers_proto_rawDescData
}

var file_external_infra_proxy_request_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_external_infra_proxy_request_servers_proto_goTypes = []interface{}{
	(*CreateServer)(nil),     // 0: chef.automate.api.infra_proxy.request.CreateServer
	(*UpdateServer)(nil),     // 1: chef.automate.api.infra_proxy.request.UpdateServer
	(*DeleteServer)(nil),     // 2: chef.automate.api.infra_proxy.request.DeleteServer
	(*GetServers)(nil),       // 3: chef.automate.api.infra_proxy.request.GetServers
	(*GetServer)(nil),        // 4: chef.automate.api.infra_proxy.request.GetServer
	(*GetServerStatus)(nil),  // 5: chef.automate.api.infra_proxy.request.GetServerStatus
	(*ValidateWebuiKey)(nil), // 6: chef.automate.api.infra_proxy.request.ValidateWebuiKey
}
var file_external_infra_proxy_request_servers_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_request_servers_proto_init() }
func file_external_infra_proxy_request_servers_proto_init() {
	if File_external_infra_proxy_request_servers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_request_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_infra_proxy_request_servers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_infra_proxy_request_servers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_request_servers_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_request_servers_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_request_servers_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_request_servers_proto = out.File
	file_external_infra_proxy_request_servers_proto_rawDesc = nil
	file_external_infra_proxy_request_servers_proto_goTypes = nil
	file_external_infra_proxy_request_servers_proto_depIdxs = nil
}
