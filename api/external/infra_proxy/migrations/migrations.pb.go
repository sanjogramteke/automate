// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: external/infra_proxy/migrations/migrations.proto

package migrations

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	request "github.com/chef/automate/api/external/infra_proxy/migrations/request"
	response "github.com/chef/automate/api/external/infra_proxy/migrations/response"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_external_infra_proxy_migrations_migrations_proto protoreflect.FileDescriptor

var file_external_infra_proxy_migrations_migrations_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x28, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x39, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x0a, 0x0a, 0x1a,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4, 0x02, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x7b,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x8a, 0xb5, 0x18,
	0x14, 0x0a, 0x12, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x67, 0x65,
	0x74, 0x12, 0xd8, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x4a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xae, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x4e, 0x12, 0x4c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x8a, 0xb5, 0x18, 0x3a, 0x0a, 0x38, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x7b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x7b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d,
	0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x12, 0x9a, 0x02, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x46,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x12, 0x3b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x30, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x7b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x8a, 0xb5, 0x18, 0x14, 0x0a, 0x12, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x8a, 0xb5, 0x18,
	0x18, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x12, 0xce, 0x02, 0x0a, 0x0e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x40, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x41,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x22, 0xb6, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x22, 0x4b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x7b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a,
	0x8a, 0xb5, 0x18, 0x3a, 0x0a, 0x38, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a,
	0x7b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x8a, 0xb5,
	0x18, 0x18, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_external_infra_proxy_migrations_migrations_proto_goTypes = []interface{}{
	(*request.GetMigrationStatusRequest)(nil),   // 0: chef.automate.api.infra_proxy.migrations.request.GetMigrationStatusRequest
	(*request.CancelMigrationRequest)(nil),      // 1: chef.automate.api.infra_proxy.migrations.request.CancelMigrationRequest
	(*request.GetStagedDataRequest)(nil),        // 2: chef.automate.api.infra_proxy.migrations.request.GetStagedDataRequest
	(*request.ConfirmPreview)(nil),              // 3: chef.automate.api.infra_proxy.migrations.request.ConfirmPreview
	(*response.GetMigrationStatusResponse)(nil), // 4: chef.automate.api.infra_proxy.migrations.response.GetMigrationStatusResponse
	(*response.CancelMigrationResponse)(nil),    // 5: chef.automate.api.infra_proxy.migrations.response.CancelMigrationResponse
	(*response.GetStagedDataResponse)(nil),      // 6: chef.automate.api.infra_proxy.migrations.response.GetStagedDataResponse
	(*response.ConfirmPreview)(nil),             // 7: chef.automate.api.infra_proxy.migrations.response.ConfirmPreview
}
var file_external_infra_proxy_migrations_migrations_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.GetMigrationStatus:input_type -> chef.automate.api.infra_proxy.migrations.request.GetMigrationStatusRequest
	1, // 1: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.CancelMigration:input_type -> chef.automate.api.infra_proxy.migrations.request.CancelMigrationRequest
	2, // 2: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.GetStagedData:input_type -> chef.automate.api.infra_proxy.migrations.request.GetStagedDataRequest
	3, // 3: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.ConfirmPreview:input_type -> chef.automate.api.infra_proxy.migrations.request.ConfirmPreview
	4, // 4: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.GetMigrationStatus:output_type -> chef.automate.api.infra_proxy.migrations.response.GetMigrationStatusResponse
	5, // 5: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.CancelMigration:output_type -> chef.automate.api.infra_proxy.migrations.response.CancelMigrationResponse
	6, // 6: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.GetStagedData:output_type -> chef.automate.api.infra_proxy.migrations.response.GetStagedDataResponse
	7, // 7: chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService.ConfirmPreview:output_type -> chef.automate.api.infra_proxy.migrations.response.ConfirmPreview
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_migrations_migrations_proto_init() }
func file_external_infra_proxy_migrations_migrations_proto_init() {
	if File_external_infra_proxy_migrations_migrations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_infra_proxy_migrations_migrations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_infra_proxy_migrations_migrations_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_migrations_migrations_proto_depIdxs,
	}.Build()
	File_external_infra_proxy_migrations_migrations_proto = out.File
	file_external_infra_proxy_migrations_migrations_proto_rawDesc = nil
	file_external_infra_proxy_migrations_migrations_proto_goTypes = nil
	file_external_infra_proxy_migrations_migrations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InfraProxyMigrationServiceClient is the client API for InfraProxyMigrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfraProxyMigrationServiceClient interface {
	GetMigrationStatus(ctx context.Context, in *request.GetMigrationStatusRequest, opts ...grpc.CallOption) (*response.GetMigrationStatusResponse, error)
	CancelMigration(ctx context.Context, in *request.CancelMigrationRequest, opts ...grpc.CallOption) (*response.CancelMigrationResponse, error)
	GetStagedData(ctx context.Context, in *request.GetStagedDataRequest, opts ...grpc.CallOption) (*response.GetStagedDataResponse, error)
	ConfirmPreview(ctx context.Context, in *request.ConfirmPreview, opts ...grpc.CallOption) (*response.ConfirmPreview, error)
}

type infraProxyMigrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfraProxyMigrationServiceClient(cc grpc.ClientConnInterface) InfraProxyMigrationServiceClient {
	return &infraProxyMigrationServiceClient{cc}
}

func (c *infraProxyMigrationServiceClient) GetMigrationStatus(ctx context.Context, in *request.GetMigrationStatusRequest, opts ...grpc.CallOption) (*response.GetMigrationStatusResponse, error) {
	out := new(response.GetMigrationStatusResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/GetMigrationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyMigrationServiceClient) CancelMigration(ctx context.Context, in *request.CancelMigrationRequest, opts ...grpc.CallOption) (*response.CancelMigrationResponse, error) {
	out := new(response.CancelMigrationResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/CancelMigration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyMigrationServiceClient) GetStagedData(ctx context.Context, in *request.GetStagedDataRequest, opts ...grpc.CallOption) (*response.GetStagedDataResponse, error) {
	out := new(response.GetStagedDataResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/GetStagedData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyMigrationServiceClient) ConfirmPreview(ctx context.Context, in *request.ConfirmPreview, opts ...grpc.CallOption) (*response.ConfirmPreview, error) {
	out := new(response.ConfirmPreview)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/ConfirmPreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraProxyMigrationServiceServer is the server API for InfraProxyMigrationService service.
type InfraProxyMigrationServiceServer interface {
	GetMigrationStatus(context.Context, *request.GetMigrationStatusRequest) (*response.GetMigrationStatusResponse, error)
	CancelMigration(context.Context, *request.CancelMigrationRequest) (*response.CancelMigrationResponse, error)
	GetStagedData(context.Context, *request.GetStagedDataRequest) (*response.GetStagedDataResponse, error)
	ConfirmPreview(context.Context, *request.ConfirmPreview) (*response.ConfirmPreview, error)
}

// UnimplementedInfraProxyMigrationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInfraProxyMigrationServiceServer struct {
}

func (*UnimplementedInfraProxyMigrationServiceServer) GetMigrationStatus(context.Context, *request.GetMigrationStatusRequest) (*response.GetMigrationStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMigrationStatus not implemented")
}
func (*UnimplementedInfraProxyMigrationServiceServer) CancelMigration(context.Context, *request.CancelMigrationRequest) (*response.CancelMigrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMigration not implemented")
}
func (*UnimplementedInfraProxyMigrationServiceServer) GetStagedData(context.Context, *request.GetStagedDataRequest) (*response.GetStagedDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStagedData not implemented")
}
func (*UnimplementedInfraProxyMigrationServiceServer) ConfirmPreview(context.Context, *request.ConfirmPreview) (*response.ConfirmPreview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPreview not implemented")
}

func RegisterInfraProxyMigrationServiceServer(s *grpc.Server, srv InfraProxyMigrationServiceServer) {
	s.RegisterService(&_InfraProxyMigrationService_serviceDesc, srv)
}

func _InfraProxyMigrationService_GetMigrationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetMigrationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyMigrationServiceServer).GetMigrationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/GetMigrationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyMigrationServiceServer).GetMigrationStatus(ctx, req.(*request.GetMigrationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxyMigrationService_CancelMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CancelMigrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyMigrationServiceServer).CancelMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/CancelMigration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyMigrationServiceServer).CancelMigration(ctx, req.(*request.CancelMigrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxyMigrationService_GetStagedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetStagedDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyMigrationServiceServer).GetStagedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/GetStagedData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyMigrationServiceServer).GetStagedData(ctx, req.(*request.GetStagedDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxyMigrationService_ConfirmPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ConfirmPreview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyMigrationServiceServer).ConfirmPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService/ConfirmPreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyMigrationServiceServer).ConfirmPreview(ctx, req.(*request.ConfirmPreview))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfraProxyMigrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.infra_proxy.migrations.InfraProxyMigrationService",
	HandlerType: (*InfraProxyMigrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMigrationStatus",
			Handler:    _InfraProxyMigrationService_GetMigrationStatus_Handler,
		},
		{
			MethodName: "CancelMigration",
			Handler:    _InfraProxyMigrationService_CancelMigration_Handler,
		},
		{
			MethodName: "GetStagedData",
			Handler:    _InfraProxyMigrationService_GetStagedData_Handler,
		},
		{
			MethodName: "ConfirmPreview",
			Handler:    _InfraProxyMigrationService_ConfirmPreview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/infra_proxy/migrations/migrations.proto",
}
