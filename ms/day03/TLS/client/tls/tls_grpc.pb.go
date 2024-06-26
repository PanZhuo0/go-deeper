// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: tls.proto

package tls

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
	T_THello_FullMethodName = "/T/THello"
)

// TClient is the client API for T service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TClient interface {
	THello(ctx context.Context, in *TReq, opts ...grpc.CallOption) (*TResp, error)
}

type tClient struct {
	cc grpc.ClientConnInterface
}

func NewTClient(cc grpc.ClientConnInterface) TClient {
	return &tClient{cc}
}

func (c *tClient) THello(ctx context.Context, in *TReq, opts ...grpc.CallOption) (*TResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TResp)
	err := c.cc.Invoke(ctx, T_THello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TServer is the server API for T service.
// All implementations must embed UnimplementedTServer
// for forward compatibility
type TServer interface {
	THello(context.Context, *TReq) (*TResp, error)
	mustEmbedUnimplementedTServer()
}

// UnimplementedTServer must be embedded to have forward compatible implementations.
type UnimplementedTServer struct {
}

func (UnimplementedTServer) THello(context.Context, *TReq) (*TResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method THello not implemented")
}
func (UnimplementedTServer) mustEmbedUnimplementedTServer() {}

// UnsafeTServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TServer will
// result in compilation errors.
type UnsafeTServer interface {
	mustEmbedUnimplementedTServer()
}

func RegisterTServer(s grpc.ServiceRegistrar, srv TServer) {
	s.RegisterService(&T_ServiceDesc, srv)
}

func _T_THello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TServer).THello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: T_THello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TServer).THello(ctx, req.(*TReq))
	}
	return interceptor(ctx, in, info, handler)
}

// T_ServiceDesc is the grpc.ServiceDesc for T service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var T_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "T",
	HandlerType: (*TServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "THello",
			Handler:    _T_THello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tls.proto",
}
