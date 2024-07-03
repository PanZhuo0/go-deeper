package main

import (
	"bi/server/bi"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1.new grpc client
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. new pb client
	c := bi.NewGretterClient(conn)
	// 3.获取流对象
	stream, err := c.BiTest(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 4.将数据通过流发送出去,一次发送一个
	names := []string{`张三`, `李四`, `王五`, `赵六`}
	for _, name := range names {
		fmt.Println("send request to server...")
		err := stream.Send(&bi.Request{
			Name: name,
		})
		if err != nil {
			log.Fatal(err)
			stream.CloseSend() //客户端出错了，需要即可把流给关闭，否则服务端会一直等待
			return
		}
		// 从stream中获取响应
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
			stream.CloseSend() //服务端出错,也需要把流给关闭,确保安全
			return
		}
		fmt.Println("got response from server:", resp.Reply)
	}
	// 所有请求发送完毕,结束流
	/* 更高阶的方法:通过信号量控制请求的退出,结束流的时机*/
	err = stream.CloseSend()
	if err != nil {
		// 关闭流都出错了怎么办?
		log.Fatal(err)
	}
}
