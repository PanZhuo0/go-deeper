package main

import (
	"context"
	"demo/add/server/add"
	"log"
	"net"

	"google.golang.org/grpc"
)

/* server端 */
type server struct {
	add.UnimplementedAdderServer
}

func main() {
	// 1.设置grpc server 的监听端口
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. new grpc server
	s := grpc.NewServer()
	// 3. 把pb服务注册到grpc server上
	add.RegisterAdderServer(s, &server{})
	// 4.开启grpc监听
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 实现服务
func (s *server) Add(ctx context.Context, req *add.AddRequest) (resp *add.AddResponse, err error) {
	return &add.AddResponse{
		Sum: int64(req.GetX()) + int64(req.GetY()),
	}, nil
}
