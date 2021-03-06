// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: average/protocolbuffer/average.proto

package protocolbuffer

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

type ComputeAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ComputeAverage) Reset() {
	*x = ComputeAverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_protocolbuffer_average_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverage) ProtoMessage() {}

func (x *ComputeAverage) ProtoReflect() protoreflect.Message {
	mi := &file_average_protocolbuffer_average_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverage.ProtoReflect.Descriptor instead.
func (*ComputeAverage) Descriptor() ([]byte, []int) {
	return file_average_protocolbuffer_average_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeAverage) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAverageStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputeAverage *ComputeAverage `protobuf:"bytes,1,opt,name=computeAverage,proto3" json:"computeAverage,omitempty"`
}

func (x *ComputeAverageStreamRequest) Reset() {
	*x = ComputeAverageStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_protocolbuffer_average_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageStreamRequest) ProtoMessage() {}

func (x *ComputeAverageStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_average_protocolbuffer_average_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageStreamRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageStreamRequest) Descriptor() ([]byte, []int) {
	return file_average_protocolbuffer_average_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeAverageStreamRequest) GetComputeAverage() *ComputeAverage {
	if x != nil {
		return x.ComputeAverage
	}
	return nil
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_protocolbuffer_average_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_average_protocolbuffer_average_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_average_protocolbuffer_average_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeAverageResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

var File_average_protocolbuffer_average_proto protoreflect.FileDescriptor

var file_average_protocolbuffer_average_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x28, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x1b, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x32, 0x74, 0x0a,
	0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_average_protocolbuffer_average_proto_rawDescOnce sync.Once
	file_average_protocolbuffer_average_proto_rawDescData = file_average_protocolbuffer_average_proto_rawDesc
)

func file_average_protocolbuffer_average_proto_rawDescGZIP() []byte {
	file_average_protocolbuffer_average_proto_rawDescOnce.Do(func() {
		file_average_protocolbuffer_average_proto_rawDescData = protoimpl.X.CompressGZIP(file_average_protocolbuffer_average_proto_rawDescData)
	})
	return file_average_protocolbuffer_average_proto_rawDescData
}

var file_average_protocolbuffer_average_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_average_protocolbuffer_average_proto_goTypes = []interface{}{
	(*ComputeAverage)(nil),              // 0: average.ComputeAverage
	(*ComputeAverageStreamRequest)(nil), // 1: average.ComputeAverageStreamRequest
	(*ComputeAverageResponse)(nil),      // 2: average.ComputeAverageResponse
}
var file_average_protocolbuffer_average_proto_depIdxs = []int32{
	0, // 0: average.ComputeAverageStreamRequest.computeAverage:type_name -> average.ComputeAverage
	1, // 1: average.ComputeAverageService.ComputeAverage:input_type -> average.ComputeAverageStreamRequest
	2, // 2: average.ComputeAverageService.ComputeAverage:output_type -> average.ComputeAverageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_average_protocolbuffer_average_proto_init() }
func file_average_protocolbuffer_average_proto_init() {
	if File_average_protocolbuffer_average_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_average_protocolbuffer_average_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverage); i {
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
		file_average_protocolbuffer_average_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageStreamRequest); i {
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
		file_average_protocolbuffer_average_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
			RawDescriptor: file_average_protocolbuffer_average_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_average_protocolbuffer_average_proto_goTypes,
		DependencyIndexes: file_average_protocolbuffer_average_proto_depIdxs,
		MessageInfos:      file_average_protocolbuffer_average_proto_msgTypes,
	}.Build()
	File_average_protocolbuffer_average_proto = out.File
	file_average_protocolbuffer_average_proto_rawDesc = nil
	file_average_protocolbuffer_average_proto_goTypes = nil
	file_average_protocolbuffer_average_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComputeAverageServiceClient is the client API for ComputeAverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComputeAverageServiceClient interface {
	//Server Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (ComputeAverageService_ComputeAverageClient, error)
}

type computeAverageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeAverageServiceClient(cc grpc.ClientConnInterface) ComputeAverageServiceClient {
	return &computeAverageServiceClient{cc}
}

func (c *computeAverageServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (ComputeAverageService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ComputeAverageService_serviceDesc.Streams[0], "/average.ComputeAverageService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &computeAverageServiceComputeAverageClient{stream}
	return x, nil
}

type ComputeAverageService_ComputeAverageClient interface {
	Send(*ComputeAverageStreamRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type computeAverageServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *computeAverageServiceComputeAverageClient) Send(m *ComputeAverageStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *computeAverageServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComputeAverageServiceServer is the server API for ComputeAverageService service.
type ComputeAverageServiceServer interface {
	//Server Streaming
	ComputeAverage(ComputeAverageService_ComputeAverageServer) error
}

// UnimplementedComputeAverageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedComputeAverageServiceServer struct {
}

func (*UnimplementedComputeAverageServiceServer) ComputeAverage(ComputeAverageService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterComputeAverageServiceServer(s *grpc.Server, srv ComputeAverageServiceServer) {
	s.RegisterService(&_ComputeAverageService_serviceDesc, srv)
}

func _ComputeAverageService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComputeAverageServiceServer).ComputeAverage(&computeAverageServiceComputeAverageServer{stream})
}

type ComputeAverageService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageStreamRequest, error)
	grpc.ServerStream
}

type computeAverageServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *computeAverageServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *computeAverageServiceComputeAverageServer) Recv() (*ComputeAverageStreamRequest, error) {
	m := new(ComputeAverageStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ComputeAverageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "average.ComputeAverageService",
	HandlerType: (*ComputeAverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeAverage",
			Handler:       _ComputeAverageService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "average/protocolbuffer/average.proto",
}
