package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"rev/rev"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	rev.UnimplementedTestServer
}

func (s *server) SayHello(ctx context.Context, in *rev.Request) (res *rev.Response, err error) {
	return &rev.Response{Reply: "你好," + in.GetName()}, nil
}

func main() {
	// 1.listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	// 2.grpc server
	s := grpc.NewServer()
	// 3.register grpc
	rev.RegisterTestServer(s, &server{})
	// 4.run
	go gw()
	err = s.Serve(l)
	if err != nil {
		fmt.Println(err)
	}
}

func gw() {
	// HTTP 连接依赖于grpc client
	// grpc new grpc client
	conn, err := grpc.NewClient(`localhost:8080`,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
	}

	// 使用runtime.NewServerMux()
	gwMux := runtime.NewServeMux()
	err = rev.RegisterTestHandler(
		context.Background(),
		gwMux,
		conn)

	if err != nil {
		fmt.Println(err)
	}

	err = http.ListenAndServe(":8081", gwMux)
	if err != nil {
		fmt.Println(err)
	}

}
