// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/echoService.proto

package echoService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetEchoMessage struct {
	TargetEcho           string   `protobuf:"bytes,1,opt,name=target_echo,json=targetEcho,proto3" json:"target_echo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEchoMessage) Reset()         { *m = GetEchoMessage{} }
func (m *GetEchoMessage) String() string { return proto.CompactTextString(m) }
func (*GetEchoMessage) ProtoMessage()    {}
func (*GetEchoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7942b665fd1403f0, []int{0}
}

func (m *GetEchoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEchoMessage.Unmarshal(m, b)
}
func (m *GetEchoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEchoMessage.Marshal(b, m, deterministic)
}
func (m *GetEchoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEchoMessage.Merge(m, src)
}
func (m *GetEchoMessage) XXX_Size() int {
	return xxx_messageInfo_GetEchoMessage.Size(m)
}
func (m *GetEchoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEchoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetEchoMessage proto.InternalMessageInfo

func (m *GetEchoMessage) GetTargetEcho() string {
	if m != nil {
		return m.TargetEcho
	}
	return ""
}

type EchoResponse struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7942b665fd1403f0, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func init() {
	proto.RegisterType((*GetEchoMessage)(nil), "GetEchoMessage")
	proto.RegisterType((*EchoResponse)(nil), "EchoResponse")
}

func init() { proto.RegisterFile("service/echoService.proto", fileDescriptor_7942b665fd1403f0) }

var fileDescriptor_7942b665fd1403f0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x4d, 0xce, 0xc8, 0x0f, 0x86, 0xb0, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x95, 0x0c, 0xb9, 0xf8, 0xdc, 0x53, 0x4b, 0x5c, 0x93, 0x33, 0xf2, 0x7d, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x85, 0xe4, 0xb9, 0xb8, 0x4b, 0x12, 0x8b, 0xd2, 0x53, 0x4b, 0xe2, 0x41, 0xaa,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb8, 0x20, 0x42, 0x20, 0x75, 0x4a, 0x2a, 0x5c, 0x3c,
	0x20, 0x3a, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0x35, 0x33, 0xaf,
	0xa0, 0xb4, 0x04, 0xaa, 0x14, 0xc2, 0x31, 0x32, 0xe6, 0x62, 0x01, 0xa9, 0x12, 0xd2, 0xe6, 0x62,
	0x87, 0x5a, 0x20, 0xc4, 0xaf, 0x87, 0x6a, 0x95, 0x14, 0xaf, 0x1e, 0xb2, 0x41, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0x47, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x8b, 0x22, 0xde, 0xb1, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	GetEcho(ctx context.Context, in *GetEchoMessage, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) GetEcho(ctx context.Context, in *GetEchoMessage, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/Echo/GetEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	GetEcho(context.Context, *GetEchoMessage) (*EchoResponse, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) GetEcho(ctx context.Context, req *GetEchoMessage) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcho not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_GetEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).GetEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Echo/GetEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).GetEcho(ctx, req.(*GetEchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEcho",
			Handler:    _Echo_GetEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/echoService.proto",
}