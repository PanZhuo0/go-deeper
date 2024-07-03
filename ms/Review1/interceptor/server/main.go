package main

import (
	"interceptor/client/interceptor"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	interceptor.UnimplementedGreeterServer
}

/* 拦截器服务端实现*/
func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	interceptor.RegisterGreeterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}
