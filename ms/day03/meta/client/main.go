package main

import (
	"context"
	"fmt"
	"meta/client/meta"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// send
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// new grpc client failed
		panic(err)
	}
	defer conn.Close()
	c := meta.NewTestClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// add metadata 2 ctx (just like header of HTTP)
	// meta token for authrization
	md := metadata.Pairs(
		"token", "app-test-zhangsan",
	)
	fmt.Printf("md:%#v\n", md)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// say hello
	var header, trailer metadata.MD // wait for response from server
	resp, err := c.SayHello(ctx, &meta.HelloRequest{Name: "name_张三"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response header:", header)
	fmt.Println(resp.Reply)
	fmt.Println("response trailer:", trailer)
}
