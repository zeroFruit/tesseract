// Code generated by protoc-gen-go. DO NOT EDIT.
// source: default_service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	default_service.proto

It has these top-level messages:
	Request
	Response
*/
package pb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Tx []byte `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type Response struct {
	Result string `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=Method" json:"Method,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Error  string `protobuf:"bytes,4,opt,name=Error" json:"Error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DefaultService service

type DefaultServiceClient interface {
	RunICode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type defaultServiceClient struct {
	cc *grpc.ClientConn
}

func NewDefaultServiceClient(cc *grpc.ClientConn) DefaultServiceClient {
	return &defaultServiceClient{cc}
}

func (c *defaultServiceClient) RunICode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pb.DefaultService/RunICode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DefaultService service

type DefaultServiceServer interface {
	RunICode(context.Context, *Request) (*Response, error)
}

func RegisterDefaultServiceServer(s *grpc.Server, srv DefaultServiceServer) {
	s.RegisterService(&_DefaultService_serviceDesc, srv)
}

func _DefaultService_RunICode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).RunICode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DefaultService/RunICode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).RunICode(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _DefaultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DefaultService",
	HandlerType: (*DefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunICode",
			Handler:    _DefaultService_RunICode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "default_service.proto",
}

func init() { proto.RegisterFile("default_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x49, 0x4d, 0x4b,
	0x2c, 0xcd, 0x29, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe4, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0xe2, 0xe3, 0x62, 0x0a, 0xa9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0x0a, 0xa9,
	0x50, 0x4a, 0xe1, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62,
	0x0b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x01, 0xcb, 0x73, 0x06, 0x41, 0x79, 0x20, 0x71, 0xdf, 0xd4,
	0x92, 0x8c, 0xfc, 0x14, 0x09, 0x26, 0x88, 0x38, 0x84, 0x27, 0x24, 0xc4, 0xc5, 0xe2, 0x92, 0x58,
	0x92, 0x28, 0xc1, 0x0c, 0x36, 0x0d, 0xcc, 0x16, 0x12, 0xe1, 0x62, 0x75, 0x2d, 0x2a, 0xca, 0x2f,
	0x92, 0x60, 0x01, 0x2b, 0x85, 0x70, 0x8c, 0x2c, 0xb9, 0xf8, 0x5c, 0x20, 0xae, 0x0b, 0x86, 0x38,
	0x4e, 0x48, 0x9d, 0x8b, 0x23, 0xa8, 0x34, 0xcf, 0xd3, 0x39, 0x3f, 0x25, 0x55, 0x88, 0x5b, 0xaf,
	0x20, 0x49, 0x0f, 0xea, 0x40, 0x29, 0x1e, 0x08, 0x07, 0xe2, 0x24, 0x25, 0x86, 0x24, 0x36, 0xb0,
	0x37, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x6f, 0x6a, 0xab, 0xdf, 0x00, 0x00, 0x00,
}