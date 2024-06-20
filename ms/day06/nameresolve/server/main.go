package main

import (
	"context"
	"fmt"
	"lb/client/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (res *pb.HelloResponse, err error) {
	return &pb.HelloResponse{Reply: "你好," + in.GetName()}, nil
}

func main() {
	// 1.listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("new listen failed,err:", err)
		return
	}
	// 2.grpc server
	s := grpc.NewServer()
	// 3.register pb
	pb.RegisterGreeterServer(s, &server{})
	// 4.开工
	err = s.Serve(l)
	if err != nil {
		fmt.Println(err)
		return
	}
}
