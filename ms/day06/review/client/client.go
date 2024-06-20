package main

import (
	"context"
	"fmt"
	"rev/rev"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	c := rev.NewTestClient(conn)
	resp, err := c.SayHello(context.Background(), &rev.Request{Name: "张三"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Reply)
}
