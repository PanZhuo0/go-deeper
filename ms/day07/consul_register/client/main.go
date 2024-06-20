package main

import (
	"c/server/c"
	"context"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1.new grpc client
	//直接使用consul解析器即可
	conn, err := grpc.NewClient(`consul://localhost:8500/hello?wait=14s`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.NewClient failed,err:", err)
		return
	}
	// 2.pb new client
	client := c.NewConsulClient(conn)
	reply, err := client.SayHello(context.Background(), &c.Request{Name: "张三"})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(reply.Reply)
}
