package main

import (
	"context"
	"fmt"
	"meta/server/meta"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	meta.UnimplementedTestServer
}

func main() {
	// listen
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	// new grpc server
	s := grpc.NewServer()
	meta.RegisterTestServer(s, &server{}) //register sever 2 grpc

	// start server
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}

func (s *server) SayHello(ctx context.Context, r *meta.HelloRequest) (resp *meta.HelloResponse, err error) {
	// trailer 2 context
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()

	// get toen from request
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println(err)
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}

	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-zhangsan" {
		return nil, status.Error(codes.Unauthenticated, "无效的token")
	}

	if vl[0] == "app-test-zhangsan" {
		// add metadata(header) to context
		header := metadata.New(map[string]string{
			"location": "广州",
		})
		grpc.SendHeader(ctx, header)
		// response
		resp = &meta.HelloResponse{Reply: "你好," + r.GetName()}
		return resp, nil
	} else {
		return nil, status.Error(codes.PermissionDenied, "token 有误")
	}
}
