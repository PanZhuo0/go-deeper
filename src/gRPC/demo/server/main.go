package main

import (
	"context"
	"fmt"
	"net"

	// 导入protoc产生的文件
	pb "gRPC/demo/server/proto"

	"google.golang.org/grpc"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	s := grpc.NewServer()
	// grpc服务端注册对应的服务
	pb.RegisterSayHelloServer(s, &server{})
	// 启动服务:将在本地9090运行
	err := s.Serve(listen)
	if err != nil {
		fmt.Println("服务启动失败")
	}
}
