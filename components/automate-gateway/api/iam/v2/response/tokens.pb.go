// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/response/tokens.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateTokenResp) Reset()         { *m = CreateTokenResp{} }
func (m *CreateTokenResp) String() string { return proto.CompactTextString(m) }
func (*CreateTokenResp) ProtoMessage()    {}
func (*CreateTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbfa53120193436, []int{0}
}

func (m *CreateTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenResp.Unmarshal(m, b)
}
func (m *CreateTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenResp.Marshal(b, m, deterministic)
}
func (m *CreateTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenResp.Merge(m, src)
}
func (m *CreateTokenResp) XXX_Size() int {
	return xxx_messageInfo_CreateTokenResp.Size(m)
}
func (m *CreateTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenResp proto.InternalMessageInfo

func (m *CreateTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetTokenResp) Reset()         { *m = GetTokenResp{} }
func (m *GetTokenResp) String() string { return proto.CompactTextString(m) }
func (*GetTokenResp) ProtoMessage()    {}
func (*GetTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbfa53120193436, []int{1}
}

func (m *GetTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResp.Unmarshal(m, b)
}
func (m *GetTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResp.Marshal(b, m, deterministic)
}
func (m *GetTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResp.Merge(m, src)
}
func (m *GetTokenResp) XXX_Size() int {
	return xxx_messageInfo_GetTokenResp.Size(m)
}
func (m *GetTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResp proto.InternalMessageInfo

func (m *GetTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type UpdateTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateTokenResp) Reset()         { *m = UpdateTokenResp{} }
func (m *UpdateTokenResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenResp) ProtoMessage()    {}
func (*UpdateTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbfa53120193436, []int{2}
}

func (m *UpdateTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenResp.Unmarshal(m, b)
}
func (m *UpdateTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenResp.Marshal(b, m, deterministic)
}
func (m *UpdateTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenResp.Merge(m, src)
}
func (m *UpdateTokenResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenResp.Size(m)
}
func (m *UpdateTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenResp proto.InternalMessageInfo

func (m *UpdateTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type DeleteTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResp) Reset()         { *m = DeleteTokenResp{} }
func (m *DeleteTokenResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResp) ProtoMessage()    {}
func (*DeleteTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbfa53120193436, []int{3}
}

func (m *DeleteTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResp.Unmarshal(m, b)
}
func (m *DeleteTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResp.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResp.Merge(m, src)
}
func (m *DeleteTokenResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResp.Size(m)
}
func (m *DeleteTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResp proto.InternalMessageInfo

type ListTokensResp struct {
	Tokens               []*common.Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListTokensResp) Reset()         { *m = ListTokensResp{} }
func (m *ListTokensResp) String() string { return proto.CompactTextString(m) }
func (*ListTokensResp) ProtoMessage()    {}
func (*ListTokensResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbfa53120193436, []int{4}
}

func (m *ListTokensResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTokensResp.Unmarshal(m, b)
}
func (m *ListTokensResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTokensResp.Marshal(b, m, deterministic)
}
func (m *ListTokensResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTokensResp.Merge(m, src)
}
func (m *ListTokensResp) XXX_Size() int {
	return xxx_messageInfo_ListTokensResp.Size(m)
}
func (m *ListTokensResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTokensResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListTokensResp proto.InternalMessageInfo

func (m *ListTokensResp) GetTokens() []*common.Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTokenResp)(nil), "chef.automate.api.iam.v2.CreateTokenResp")
	proto.RegisterType((*GetTokenResp)(nil), "chef.automate.api.iam.v2.GetTokenResp")
	proto.RegisterType((*UpdateTokenResp)(nil), "chef.automate.api.iam.v2.UpdateTokenResp")
	proto.RegisterType((*DeleteTokenResp)(nil), "chef.automate.api.iam.v2.DeleteTokenResp")
	proto.RegisterType((*ListTokensResp)(nil), "chef.automate.api.iam.v2.ListTokensResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/response/tokens.proto", fileDescriptor_ddbfa53120193436)
}

var fileDescriptor_ddbfa53120193436 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x4d, 0x8b, 0x14, 0x31,
	0x10, 0x25, 0x8a, 0x7b, 0x68, 0xc5, 0xc5, 0x39, 0x0d, 0x7b, 0x31, 0xe4, 0xe4, 0xa1, 0x3b, 0xc1,
	0x16, 0x11, 0x1a, 0x2f, 0xab, 0x82, 0x22, 0x9e, 0x44, 0x2f, 0xca, 0x80, 0x35, 0x3d, 0xb5, 0x3d,
	0xd1, 0x4d, 0x2a, 0x24, 0xd5, 0x33, 0xc8, 0x32, 0x7f, 0x62, 0xff, 0x85, 0x17, 0xc1, 0xdf, 0x22,
	0x08, 0x1e, 0xfc, 0x2f, 0xd2, 0x69, 0xdd, 0x2f, 0x11, 0xd6, 0x83, 0x87, 0xb9, 0xa5, 0x1e, 0xaf,
	0xf2, 0x5e, 0x3d, 0x52, 0x29, 0x1e, 0xb6, 0xe4, 0x02, 0x79, 0xf4, 0x9c, 0x0c, 0xf4, 0x4c, 0x0e,
	0x18, 0xab, 0x0e, 0x18, 0xd7, 0xf0, 0xd1, 0x40, 0xb0, 0xc6, 0x82, 0x33, 0xab, 0xda, 0x44, 0x4c,
	0x81, 0x7c, 0x42, 0xc3, 0xf4, 0x01, 0x7d, 0xd2, 0x21, 0x12, 0xd3, 0x64, 0xda, 0x2e, 0xf1, 0x40,
	0xff, 0xee, 0xd3, 0x10, 0xac, 0xb6, 0xe0, 0xf4, 0xaa, 0xde, 0x2b, 0x33, 0xa1, 0xad, 0x3a, 0xf4,
	0x55, 0x5a, 0x43, 0xd7, 0x61, 0x34, 0x14, 0xd8, 0x92, 0x4f, 0x06, 0xbc, 0x27, 0x86, 0x7c, 0x1e,
	0xef, 0xd9, 0x6b, 0x2e, 0xe9, 0xa2, 0x25, 0xe7, 0xc8, 0x9f, 0xf3, 0xa0, 0xbe, 0x88, 0x62, 0xf7,
	0x71, 0x44, 0x60, 0x7c, 0x35, 0xc0, 0x2f, 0x31, 0x85, 0xc9, 0xfd, 0xe2, 0x5a, 0xe6, 0x4c, 0x85,
	0x14, 0x77, 0xae, 0xd7, 0xb7, 0xf5, 0xdf, 0x7c, 0xea, 0xb1, 0x67, 0x64, 0x37, 0xdd, 0xf1, 0xfe,
	0xa2, 0x9e, 0x4f, 0xde, 0x1d, 0x29, 0x0f, 0x0e, 0x55, 0x23, 0x55, 0xc6, 0xe5, 0x5d, 0x55, 0x4a,
	0x65, 0x17, 0x27, 0x40, 0x95, 0x01, 0x68, 0xd9, 0xae, 0x06, 0x16, 0xc7, 0x1e, 0x4b, 0xa9, 0x42,
	0xa4, 0xf7, 0xd8, 0x72, 0x52, 0x8d, 0x7c, 0xab, 0x10, 0x12, 0x57, 0x11, 0x3b, 0x4b, 0x7e, 0xe0,
	0xae, 0xf1, 0xb4, 0x9c, 0x6d, 0xd4, 0x67, 0x51, 0xdc, 0x78, 0x8a, 0xbc, 0x3d, 0x86, 0x87, 0x90,
	0x5f, 0x87, 0xc5, 0x56, 0x85, 0x7c, 0xab, 0xd8, 0x7d, 0x82, 0x87, 0x78, 0xc6, 0xb2, 0xfa, 0x74,
	0xa5, 0xb8, 0xf9, 0xc2, 0xa6, 0x31, 0xf8, 0x94, 0xa7, 0x78, 0x50, 0xec, 0x8c, 0xcf, 0x69, 0x2a,
	0xe4, 0xd5, 0xcb, 0x8c, 0xf1, 0x8b, 0xde, 0xfc, 0x10, 0xc7, 0xfb, 0xdf, 0x45, 0xfd, 0x4d, 0x4c,
	0xbe, 0x8a, 0xa3, 0xd1, 0x72, 0x76, 0xf5, 0xbf, 0xe7, 0x2a, 0xe5, 0x45, 0x85, 0xfa, 0xa2, 0x42,
	0x7d, 0x4e, 0xe1, 0x00, 0x0e, 0xd3, 0x1f, 0x12, 0x9e, 0x22, 0x2f, 0xcf, 0x68, 0x24, 0xea, 0x4f,
	0xeb, 0xd9, 0x66, 0xb6, 0x79, 0xf4, 0xfc, 0xcd, 0xb3, 0xce, 0xf2, 0xb2, 0x9f, 0xeb, 0x96, 0x9c,
	0x19, 0x42, 0x39, 0x59, 0x4d, 0xf3, 0x8f, 0x9f, 0xc6, 0x7c, 0x27, 0xaf, 0xea, 0xbd, 0x9f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd0, 0x22, 0x65, 0x59, 0x6e, 0x04, 0x00, 0x00,
}
