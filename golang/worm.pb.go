// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worm.proto

package golang

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Result of an rpc call
type Response struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Code    int32  `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// Subscribe for Events
type Subscription struct {
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Subscription) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Subscription)(nil), "Subscription")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WormService service

type WormServiceClient interface {
	// Issue a Command
	SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	// Subscribe, receive Events
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (WormService_SubscribeClient, error)
}

type wormServiceClient struct {
	cc *grpc.ClientConn
}

func NewWormServiceClient(cc *grpc.ClientConn) WormServiceClient {
	return &wormServiceClient{cc}
}

func (c *wormServiceClient) SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/WormService/SendCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wormServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (WormService_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WormService_serviceDesc.Streams[0], c.cc, "/WormService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &wormServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WormService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type wormServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *wormServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for WormService service

type WormServiceServer interface {
	// Issue a Command
	SendCommand(context.Context, *Command) (*Response, error)
	// Subscribe, receive Events
	Subscribe(*Subscription, WormService_SubscribeServer) error
}

func RegisterWormServiceServer(s *grpc.Server, srv WormServiceServer) {
	s.RegisterService(&_WormService_serviceDesc, srv)
}

func _WormService_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WormServiceServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WormService/SendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WormServiceServer).SendCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _WormService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WormServiceServer).Subscribe(m, &wormServiceSubscribeServer{stream})
}

type WormService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type wormServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *wormServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _WormService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WormService",
	HandlerType: (*WormServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommand",
			Handler:    _WormService_SendCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _WormService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "worm.proto",
}

func init() { proto.RegisterFile("worm.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x6b, 0xa0, 0x21, 0xb9, 0xb4, 0xcb, 0x89, 0x21, 0xca, 0x14, 0x45, 0x15, 0xca, 0x64,
	0x21, 0xf8, 0x07, 0x20, 0x36, 0x26, 0x67, 0x40, 0x82, 0x29, 0x71, 0x4e, 0x55, 0x24, 0xe2, 0xb3,
	0x6c, 0xb7, 0xfc, 0x7d, 0x84, 0xeb, 0x48, 0xdd, 0xbe, 0xef, 0x86, 0x77, 0xef, 0x01, 0xfc, 0xb2,
	0x5b, 0xa4, 0x75, 0x1c, 0xb8, 0x2e, 0xdf, 0xcf, 0x64, 0x42, 0x92, 0xfd, 0x1b, 0x2f, 0xcb, 0x60,
	0xa6, 0x8b, 0xb6, 0x1f, 0x90, 0x2b, 0xf2, 0x96, 0x8d, 0x27, 0xac, 0xe0, 0xde, 0x9f, 0xb4, 0x26,
	0xef, 0x2b, 0xd1, 0x88, 0x2e, 0x57, 0xab, 0x22, 0xc2, 0xdd, 0x34, 0x84, 0xa1, 0xba, 0x69, 0x44,
	0xb7, 0x53, 0x91, 0xff, 0x6f, 0x9a, 0x27, 0xaa, 0x6e, 0x1b, 0xd1, 0x6d, 0x55, 0xe4, 0xf6, 0x00,
	0xbb, 0xfe, 0x34, 0x7a, 0xed, 0x66, 0x1b, 0x66, 0x36, 0xf8, 0x00, 0xdb, 0xc0, 0x76, 0xd6, 0x31,
	0xaf, 0x50, 0x17, 0x79, 0xfe, 0x86, 0xf2, 0x93, 0xdd, 0xd2, 0x93, 0x3b, 0xcf, 0x9a, 0xf0, 0x00,
	0x65, 0x4f, 0x66, 0x4a, 0xbd, 0x30, 0x97, 0x89, 0xea, 0x42, 0xae, 0xd5, 0xda, 0x0d, 0x3e, 0x42,
	0x91, 0xa2, 0x47, 0xc2, 0xbd, 0xbc, 0x7e, 0x53, 0x67, 0x32, 0x2e, 0x6c, 0x37, 0x4f, 0xe2, 0x35,
	0xff, 0xca, 0x8e, 0xfc, 0x33, 0x98, 0xe3, 0x98, 0xc5, 0x85, 0x2f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x3d, 0x42, 0x11, 0x0b, 0x01, 0x00, 0x00,
}
