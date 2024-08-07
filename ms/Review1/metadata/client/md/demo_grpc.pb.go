// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: demo.proto

package md

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Gretter_SayHello_FullMethodName = "/Gretter/SayHello"
)

// GretterClient is the client API for Gretter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 这个例子中，要求MetaData携带Token信息
type GretterClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type gretterClient struct {
	cc grpc.ClientConnInterface
}

func NewGretterClient(cc grpc.ClientConnInterface) GretterClient {
	return &gretterClient{cc}
}

func (c *gretterClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Gretter_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GretterServer is the server API for Gretter service.
// All implementations must embed UnimplementedGretterServer
// for forward compatibility
//
// 这个例子中，要求MetaData携带Token信息
type GretterServer interface {
	SayHello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGretterServer()
}

// UnimplementedGretterServer must be embedded to have forward compatible implementations.
type UnimplementedGretterServer struct {
}

func (UnimplementedGretterServer) SayHello(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGretterServer) mustEmbedUnimplementedGretterServer() {}

// UnsafeGretterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GretterServer will
// result in compilation errors.
type UnsafeGretterServer interface {
	mustEmbedUnimplementedGretterServer()
}

func RegisterGretterServer(s grpc.ServiceRegistrar, srv GretterServer) {
	s.RegisterService(&Gretter_ServiceDesc, srv)
}

func _Gretter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GretterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gretter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GretterServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Gretter_ServiceDesc is the grpc.ServiceDesc for Gretter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gretter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Gretter",
	HandlerType: (*GretterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Gretter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
