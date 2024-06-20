package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sp/sameport"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
demo

	简单的一个grpc-gateway HTTP, gRPC相同端口
	居然报错了,不知道为什么出错了,先过
*/

/*
监听同一个端口:

	将gRPC请求/非gRPC请求分流给不同的Handler
	Attention:下面的代码是不完整的，仅为示例代码
*/
type server struct {
	sameport.UnimplementedSamePServer
}

func main() {
	// GRPC SERVER

	// 1.listner
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.grpc server
	s := grpc.NewServer()
	// 3.register server to pb
	sameport.RegisterSamePServer(s, &server{})
	// 4.server
	go s.Serve(l)

	// HTTP SERVER 需要使用h2c包提供的方法实现一个端口两种服务
	// 1.获取runtime Mux
	gwMux := runtime.NewServeMux()
	// 2.往pb 注册服务,from Endpoint
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = sameport.RegisterSamePHandlerFromEndpoint(
		context.Background(),
		gwMux,
		`localhost:8080`,
		dops)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3.新建HTTP mux:HTTP 多路复用器
	mux := http.NewServeMux()
	mux.Handle(`/`, gwMux)
	// 4.定义HTTP SERVER
	gwServer := &http.Server{
		Addr:    `localhost:8080`,
		Handler: grpcHandlerFunc(s, mux),
	}
	// 5.开动
	gwServer.Serve(l)
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	httphandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Context-Type"), "application/grpc") {
			//如果HTTP请求头表名请求方式是grpc
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
	return h2c.NewHandler(httphandler, &http2.Server{})
}
