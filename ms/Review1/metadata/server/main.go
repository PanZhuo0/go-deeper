package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"md/server/md"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	md.UnimplementedGretterServer
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	md.RegisterGretterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

/* 这个例子中的错误处理只是示例，仅供参考 */
func (s *server) SayHello(ctx context.Context, req *md.Request) (resp *md.Response, err error) {
	// 1.从context中取出MetaData
	md1, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// Request中没有MD
		return nil, errors.New("需要MetaData")
	}
	// 2.检验MetaData中是否存在token 这个key
	token, ok := md1["token"]
	if !ok {
		// Request有MD 不过MD中没有token 这个Key
		return nil, errors.New("MetaData中没有token")
	}
	// 3.服务端发送Header & Trailer MetaData 给Client
	header := metadata.New(
		map[string]string{
			"location": "beijing",
		},
	)
	grpc.SendHeader(ctx, header)
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", fmt.Sprint(time.Now().Unix()),
		)
		grpc.SetTrailer(ctx, trailer)
	}()

	// 4.检验token对应的值是否是预期值(这里只是示例，实际的Token更加复杂)
	if token[0] == `metadata-test` {
		resp = &md.Response{
			Reply: "你好," + req.GetName(),
		}
		return resp, nil
	} else {
		return nil, errors.New("token有误")
	}
}
