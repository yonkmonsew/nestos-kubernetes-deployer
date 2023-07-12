//
//Copyright 2023 KylinSoft  Co., Ltd.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: daemon.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UpgradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KubeVersion string `protobuf:"bytes,1,opt,name=kube_version,json=kubeVersion,proto3" json:"kube_version,omitempty"`
	OsImageUrl  string `protobuf:"bytes,2,opt,name=os_image_url,json=osImageUrl,proto3" json:"os_image_url,omitempty"`
	OsVersion   string `protobuf:"bytes,3,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
}

func (x *UpgradeRequest) Reset() {
	*x = UpgradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeRequest) ProtoMessage() {}

func (x *UpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeRequest.ProtoReflect.Descriptor instead.
func (*UpgradeRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *UpgradeRequest) GetKubeVersion() string {
	if x != nil {
		return x.KubeVersion
	}
	return ""
}

func (x *UpgradeRequest) GetOsImageUrl() string {
	if x != nil {
		return x.OsImageUrl
	}
	return ""
}

func (x *UpgradeRequest) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

type UpgradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32 `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *UpgradeResponse) Reset() {
	*x = UpgradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeResponse) ProtoMessage() {}

func (x *UpgradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeResponse.ProtoReflect.Descriptor instead.
func (*UpgradeResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{1}
}

func (x *UpgradeResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

var File_daemon_proto protoreflect.FileDescriptor

var file_daemon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x0e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x75, 0x62, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6b, 0x75, 0x62, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6f,
	0x73, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x0f,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x32, 0x4e, 0x0a, 0x0e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x16,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x25, 0x5a, 0x23, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daemon_proto_rawDescOnce sync.Once
	file_daemon_proto_rawDescData = file_daemon_proto_rawDesc
)

func file_daemon_proto_rawDescGZIP() []byte {
	file_daemon_proto_rawDescOnce.Do(func() {
		file_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_daemon_proto_rawDescData)
	})
	return file_daemon_proto_rawDescData
}

var file_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_daemon_proto_goTypes = []interface{}{
	(*UpgradeRequest)(nil),  // 0: daemon.UpgradeRequest
	(*UpgradeResponse)(nil), // 1: daemon.UpgradeResponse
}
var file_daemon_proto_depIdxs = []int32{
	0, // 0: daemon.UpgradeCluster.Upgrade:input_type -> daemon.UpgradeRequest
	1, // 1: daemon.UpgradeCluster.Upgrade:output_type -> daemon.UpgradeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_daemon_proto_init() }
func file_daemon_proto_init() {
	if File_daemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeRequest); i {
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
		file_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeResponse); i {
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
			RawDescriptor: file_daemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daemon_proto_goTypes,
		DependencyIndexes: file_daemon_proto_depIdxs,
		MessageInfos:      file_daemon_proto_msgTypes,
	}.Build()
	File_daemon_proto = out.File
	file_daemon_proto_rawDesc = nil
	file_daemon_proto_goTypes = nil
	file_daemon_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UpgradeClusterClient is the client API for UpgradeCluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpgradeClusterClient interface {
	Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error)
}

type upgradeClusterClient struct {
	cc grpc.ClientConnInterface
}

func NewUpgradeClusterClient(cc grpc.ClientConnInterface) UpgradeClusterClient {
	return &upgradeClusterClient{cc}
}

func (c *upgradeClusterClient) Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error) {
	out := new(UpgradeResponse)
	err := c.cc.Invoke(ctx, "/daemon.UpgradeCluster/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpgradeClusterServer is the server API for UpgradeCluster service.
type UpgradeClusterServer interface {
	Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error)
}

// UnimplementedUpgradeClusterServer can be embedded to have forward compatible implementations.
type UnimplementedUpgradeClusterServer struct {
}

func (*UnimplementedUpgradeClusterServer) Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upgrade not implemented")
}

func RegisterUpgradeClusterServer(s *grpc.Server, srv UpgradeClusterServer) {
	s.RegisterService(&_UpgradeCluster_serviceDesc, srv)
}

func _UpgradeCluster_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpgradeClusterServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.UpgradeCluster/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpgradeClusterServer).Upgrade(ctx, req.(*UpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpgradeCluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.UpgradeCluster",
	HandlerType: (*UpgradeClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upgrade",
			Handler:    _UpgradeCluster_Upgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon.proto",
}
