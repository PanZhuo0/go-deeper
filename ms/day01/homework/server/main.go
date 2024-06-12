package main

import (
	"context"
	"fmt"
	"net"
	"server/homework"

	"google.golang.org/grpc"
)

// 包含掉homework微服务
type server struct {
	// 可以在这增加其他字段
	homework.UnimplementedHomeWorkServer
}

// 这里应该是Add
// 如果写Sum 对应的结果会正确吗?不会，必须同名(隐式实现)
func (s *server) Add(ctx context.Context, r *homework.HomeWokrRequest) (rs *homework.HomeWorkResponse, err error) {
	sum := r.X + r.Y
	return &homework.HomeWorkResponse{Sum: sum}, nil
}

func main() {
	// 监听
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen,err:", err)
		return
	}
	// 创建grpc
	s := grpc.NewServer()                         //grpc server
	homework.RegisterHomeWorkServer(s, &server{}) //register our server to grpc server
	err = s.Serve(l)
	if err != nil {
		fmt.Println("failed to start server on listent,err:", err)
		return
	}
}
