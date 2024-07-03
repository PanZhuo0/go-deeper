package main

import (
	"context"
	"fmt"
	"log"
	"md/client/md"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	c := md.NewGretterClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	/* MD的创建有两种方法
	1.新建、然后追加
		md1 := metadata.New(nil)
		md1.Append("token", "metadata-test")
	2.Pairs
		md1 := metadata.New(nil)
		md1.Append("token", "metadata-test")
	*/
	md1 := metadata.Pairs(
		"token", "metadata-test",
	)
	/* 把Metadata对象追加到Context:
	metadata.NewOutgoingContext(ctx,md)
	*/
	ctx = metadata.NewOutgoingContext(ctx, md1)
	var header, trailer metadata.MD //声明header trailer 等待用户返回
	resp, err := c.SayHello(ctx, &md.Request{Name: "张三"},
		// 提前写入Header 和 Trailer 接收server端响应的MetaData数据
		grpc.Header(&header),
		grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 读取response
	fmt.Println("got response from server:", resp.Reply)
	// 4.client 读取从server端响应的metadata 包括Header和Trailer
	fmt.Println("got header metadata from server:", header.Get("location"))
	fmt.Println("reply body:", resp.Reply)
	fmt.Println("got trailer metadata from server:", trailer.Get("timestamp"))
}
