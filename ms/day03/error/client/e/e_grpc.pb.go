// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: e.proto

package e

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
	E_Err_FullMethodName = "/E/err"
)

// EClient is the client API for E service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EClient interface {
	Err(ctx context.Context, in *ERequest, opts ...grpc.CallOption) (*EResponse, error)
}

type eClient struct {
	cc grpc.ClientConnInterface
}

func NewEClient(cc grpc.ClientConnInterface) EClient {
	return &eClient{cc}
}

func (c *eClient) Err(ctx context.Context, in *ERequest, opts ...grpc.CallOption) (*EResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EResponse)
	err := c.cc.Invoke(ctx, E_Err_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EServer is the server API for E service.
// All implementations must embed UnimplementedEServer
// for forward compatibility
type EServer interface {
	Err(context.Context, *ERequest) (*EResponse, error)
	mustEmbedUnimplementedEServer()
}

// UnimplementedEServer must be embedded to have forward compatible implementations.
type UnimplementedEServer struct {
}

func (UnimplementedEServer) Err(context.Context, *ERequest) (*EResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Err not implemented")
}
func (UnimplementedEServer) mustEmbedUnimplementedEServer() {}

// UnsafeEServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EServer will
// result in compilation errors.
type UnsafeEServer interface {
	mustEmbedUnimplementedEServer()
}

func RegisterEServer(s grpc.ServiceRegistrar, srv EServer) {
	s.RegisterService(&E_ServiceDesc, srv)
}

func _E_Err_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EServer).Err(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: E_Err_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EServer).Err(ctx, req.(*ERequest))
	}
	return interceptor(ctx, in, info, handler)
}

// E_ServiceDesc is the grpc.ServiceDesc for E service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var E_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "E",
	HandlerType: (*EServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "err",
			Handler:    _E_Err_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "e.proto",
}
