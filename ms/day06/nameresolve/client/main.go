package main

import (
	"context"
	"fmt"
	"lb/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1.new client
	conn, err := grpc.NewClient(`qimi:///resolver.liwenzhou.com`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.register server
	c := pb.NewGreeterClient(conn)
	//
	reply, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "张三"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply.Reply)
}
