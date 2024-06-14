// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: f.proto

package flow

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
	FlowRpc_Replies_FullMethodName = "/flow_rpc/replies"
	FlowRpc_Reqs_FullMethodName    = "/flow_rpc/reqs"
	FlowRpc_Bi_FullMethodName      = "/flow_rpc/bi"
)

// FlowRpcClient is the client API for FlowRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowRpcClient interface {
	// 服务端流式
	Replies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (FlowRpc_RepliesClient, error)
	// 客户端流式
	Reqs(ctx context.Context, opts ...grpc.CallOption) (FlowRpc_ReqsClient, error)
	// 双向流
	Bi(ctx context.Context, opts ...grpc.CallOption) (FlowRpc_BiClient, error)
}

type flowRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowRpcClient(cc grpc.ClientConnInterface) FlowRpcClient {
	return &flowRpcClient{cc}
}

func (c *flowRpcClient) Replies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (FlowRpc_RepliesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowRpc_ServiceDesc.Streams[0], FlowRpc_Replies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &flowRpcRepliesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowRpc_RepliesClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type flowRpcRepliesClient struct {
	grpc.ClientStream
}

func (x *flowRpcRepliesClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowRpcClient) Reqs(ctx context.Context, opts ...grpc.CallOption) (FlowRpc_ReqsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowRpc_ServiceDesc.Streams[1], FlowRpc_Reqs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &flowRpcReqsClient{ClientStream: stream}
	return x, nil
}

type FlowRpc_ReqsClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type flowRpcReqsClient struct {
	grpc.ClientStream
}

func (x *flowRpcReqsClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowRpcReqsClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowRpcClient) Bi(ctx context.Context, opts ...grpc.CallOption) (FlowRpc_BiClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowRpc_ServiceDesc.Streams[2], FlowRpc_Bi_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &flowRpcBiClient{ClientStream: stream}
	return x, nil
}

type FlowRpc_BiClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type flowRpcBiClient struct {
	grpc.ClientStream
}

func (x *flowRpcBiClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowRpcBiClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowRpcServer is the server API for FlowRpc service.
// All implementations must embed UnimplementedFlowRpcServer
// for forward compatibility
type FlowRpcServer interface {
	// 服务端流式
	Replies(*HelloRequest, FlowRpc_RepliesServer) error
	// 客户端流式
	Reqs(FlowRpc_ReqsServer) error
	// 双向流
	Bi(FlowRpc_BiServer) error
	mustEmbedUnimplementedFlowRpcServer()
}

// UnimplementedFlowRpcServer must be embedded to have forward compatible implementations.
type UnimplementedFlowRpcServer struct {
}

func (UnimplementedFlowRpcServer) Replies(*HelloRequest, FlowRpc_RepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method Replies not implemented")
}
func (UnimplementedFlowRpcServer) Reqs(FlowRpc_ReqsServer) error {
	return status.Errorf(codes.Unimplemented, "method Reqs not implemented")
}
func (UnimplementedFlowRpcServer) Bi(FlowRpc_BiServer) error {
	return status.Errorf(codes.Unimplemented, "method Bi not implemented")
}
func (UnimplementedFlowRpcServer) mustEmbedUnimplementedFlowRpcServer() {}

// UnsafeFlowRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowRpcServer will
// result in compilation errors.
type UnsafeFlowRpcServer interface {
	mustEmbedUnimplementedFlowRpcServer()
}

func RegisterFlowRpcServer(s grpc.ServiceRegistrar, srv FlowRpcServer) {
	s.RegisterService(&FlowRpc_ServiceDesc, srv)
}

func _FlowRpc_Replies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowRpcServer).Replies(m, &flowRpcRepliesServer{ServerStream: stream})
}

type FlowRpc_RepliesServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type flowRpcRepliesServer struct {
	grpc.ServerStream
}

func (x *flowRpcRepliesServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowRpc_Reqs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowRpcServer).Reqs(&flowRpcReqsServer{ServerStream: stream})
}

type FlowRpc_ReqsServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type flowRpcReqsServer struct {
	grpc.ServerStream
}

func (x *flowRpcReqsServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowRpcReqsServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowRpc_Bi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowRpcServer).Bi(&flowRpcBiServer{ServerStream: stream})
}

type FlowRpc_BiServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type flowRpcBiServer struct {
	grpc.ServerStream
}

func (x *flowRpcBiServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowRpcBiServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowRpc_ServiceDesc is the grpc.ServiceDesc for FlowRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow_rpc",
	HandlerType: (*FlowRpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "replies",
			Handler:       _FlowRpc_Replies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "reqs",
			Handler:       _FlowRpc_Reqs_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "bi",
			Handler:       _FlowRpc_Bi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "f.proto",
}