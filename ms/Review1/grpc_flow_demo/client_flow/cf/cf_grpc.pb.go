// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: cf.proto

package cf

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
	CF_Gretter_FullMethodName = "/CF/Gretter"
)

// CFClient is the client API for CF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CFClient interface {
	// client端 流式的请求
	Gretter(ctx context.Context, opts ...grpc.CallOption) (CF_GretterClient, error)
}

type cFClient struct {
	cc grpc.ClientConnInterface
}

func NewCFClient(cc grpc.ClientConnInterface) CFClient {
	return &cFClient{cc}
}

func (c *cFClient) Gretter(ctx context.Context, opts ...grpc.CallOption) (CF_GretterClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CF_ServiceDesc.Streams[0], CF_Gretter_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &cFGretterClient{ClientStream: stream}
	return x, nil
}

type CF_GretterClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type cFGretterClient struct {
	grpc.ClientStream
}

func (x *cFGretterClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cFGretterClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CFServer is the server API for CF service.
// All implementations must embed UnimplementedCFServer
// for forward compatibility
type CFServer interface {
	// client端 流式的请求
	Gretter(CF_GretterServer) error
	mustEmbedUnimplementedCFServer()
}

// UnimplementedCFServer must be embedded to have forward compatible implementations.
type UnimplementedCFServer struct {
}

func (UnimplementedCFServer) Gretter(CF_GretterServer) error {
	return status.Errorf(codes.Unimplemented, "method Gretter not implemented")
}
func (UnimplementedCFServer) mustEmbedUnimplementedCFServer() {}

// UnsafeCFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CFServer will
// result in compilation errors.
type UnsafeCFServer interface {
	mustEmbedUnimplementedCFServer()
}

func RegisterCFServer(s grpc.ServiceRegistrar, srv CFServer) {
	s.RegisterService(&CF_ServiceDesc, srv)
}

func _CF_Gretter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CFServer).Gretter(&cFGretterServer{ServerStream: stream})
}

type CF_GretterServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type cFGretterServer struct {
	grpc.ServerStream
}

func (x *cFGretterServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cFGretterServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CF_ServiceDesc is the grpc.ServiceDesc for CF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CF",
	HandlerType: (*CFServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Gretter",
			Handler:       _CF_Gretter_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cf.proto",
}
