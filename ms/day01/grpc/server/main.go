package main

import (
	"context"
	"fmt"
	"net"
	"server/hello/pb"

	"google.golang.org/grpc"
)

// grpc server
// 要把unimplementXXXServer 在我们自己这里包起来
type server struct {
	// 可以在这里增加其他字段
	pb.UnimplementedGreeterServer //用来确保当我们没有完全实现的方法时，程序也能跑起来
}

// say hello 需要我们实现的方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	reply := "Hello," + req.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen,err:", err)
		return
	}
	// google.golang.org/grpc
	s := grpc.NewServer()                  //create grpc server
	pb.RegisterGreeterServer(s, &server{}) //register server
	err = s.Serve(l)                       // start server
	if err != nil {
		fmt.Println("failed to server,err:", err)
		return
	}
}
