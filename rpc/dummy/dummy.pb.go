// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: rpc/testing/testing.proto

package dummy

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

type DummyRpcParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *DummyRpcParam) Reset() {
	*x = DummyRpcParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_testing_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyRpcParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyRpcParam) ProtoMessage() {}

func (x *DummyRpcParam) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_testing_testing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyRpcParam.ProtoReflect.Descriptor instead.
func (*DummyRpcParam) Descriptor() ([]byte, []int) {
	return file_rpc_testing_testing_proto_rawDescGZIP(), []int{0}
}

func (x *DummyRpcParam) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *DummyRpcParam) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type DummyRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DummyRpcResponse) Reset() {
	*x = DummyRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_testing_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyRpcResponse) ProtoMessage() {}

func (x *DummyRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_testing_testing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyRpcResponse.ProtoReflect.Descriptor instead.
func (*DummyRpcResponse) Descriptor() ([]byte, []int) {
	return file_rpc_testing_testing_proto_rawDescGZIP(), []int{1}
}

func (x *DummyRpcResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_rpc_testing_testing_proto protoreflect.FileDescriptor

var file_rpc_testing_testing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x52, 0x70, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x2a, 0x0a, 0x10, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x3e, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x65, 0x72, 0x47, 0x75,
	0x69, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x70, 0x63, 0x12,
	0x0e, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x70, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x11, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_testing_testing_proto_rawDescOnce sync.Once
	file_rpc_testing_testing_proto_rawDescData = file_rpc_testing_testing_proto_rawDesc
)

func file_rpc_testing_testing_proto_rawDescGZIP() []byte {
	file_rpc_testing_testing_proto_rawDescOnce.Do(func() {
		file_rpc_testing_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_testing_testing_proto_rawDescData)
	})
	return file_rpc_testing_testing_proto_rawDescData
}

var file_rpc_testing_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_testing_testing_proto_goTypes = []interface{}{
	(*DummyRpcParam)(nil),    // 0: DummyRpcParam
	(*DummyRpcResponse)(nil), // 1: DummyRpcResponse
}
var file_rpc_testing_testing_proto_depIdxs = []int32{
	0, // 0: TesterGuide.DummyRpc:input_type -> DummyRpcParam
	1, // 1: TesterGuide.DummyRpc:output_type -> DummyRpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_testing_testing_proto_init() }
func file_rpc_testing_testing_proto_init() {
	if File_rpc_testing_testing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_testing_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyRpcParam); i {
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
		file_rpc_testing_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyRpcResponse); i {
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
			RawDescriptor: file_rpc_testing_testing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_testing_testing_proto_goTypes,
		DependencyIndexes: file_rpc_testing_testing_proto_depIdxs,
		MessageInfos:      file_rpc_testing_testing_proto_msgTypes,
	}.Build()
	File_rpc_testing_testing_proto = out.File
	file_rpc_testing_testing_proto_rawDesc = nil
	file_rpc_testing_testing_proto_goTypes = nil
	file_rpc_testing_testing_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TesterGuideClient is the client API for TesterGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TesterGuideClient interface {
	DummyRpc(ctx context.Context, in *DummyRpcParam, opts ...grpc.CallOption) (*DummyRpcResponse, error)
}

type testerGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewTesterGuideClient(cc grpc.ClientConnInterface) TesterGuideClient {
	return &testerGuideClient{cc}
}

func (c *testerGuideClient) DummyRpc(ctx context.Context, in *DummyRpcParam, opts ...grpc.CallOption) (*DummyRpcResponse, error) {
	out := new(DummyRpcResponse)
	err := c.cc.Invoke(ctx, "/TesterGuide/DummyRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TesterGuideServer is the server API for TesterGuide service.
type TesterGuideServer interface {
	DummyRpc(context.Context, *DummyRpcParam) (*DummyRpcResponse, error)
}

// UnimplementedTesterGuideServer can be embedded to have forward compatible implementations.
type UnimplementedTesterGuideServer struct {
}

func (*UnimplementedTesterGuideServer) DummyRpc(context.Context, *DummyRpcParam) (*DummyRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DummyRpc not implemented")
}

func RegisterTesterGuideServer(s *grpc.Server, srv TesterGuideServer) {
	s.RegisterService(&_TesterGuide_serviceDesc, srv)
}

func _TesterGuide_DummyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyRpcParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TesterGuideServer).DummyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TesterGuide/DummyRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TesterGuideServer).DummyRpc(ctx, req.(*DummyRpcParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _TesterGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TesterGuide",
	HandlerType: (*TesterGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DummyRpc",
			Handler:    _TesterGuide_DummyRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/testing/testing.proto",
}
