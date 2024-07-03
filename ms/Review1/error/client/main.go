package main

import (
	"context"
	"error/client/e"
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	c := e.NewGreeterClient(conn)
	resp, err := c.SayHello(context.Background(), &e.Request{
		Name: "张三",
	})
	if err != nil {
		// 解析错误中的详细信息(可能错误中携带了详细信息)
		s := status.Convert(err)
		ds := s.Details()
		for _, d := range ds {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Println("QuotaFailure: ", info)
			default:
				fmt.Println("unexpected type:", info)
			}
		}
		return
	}
	/* 客户端 */
	fmt.Println(resp.Reply)
}
