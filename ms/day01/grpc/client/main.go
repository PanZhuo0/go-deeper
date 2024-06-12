package main

import (
	"client/hello/pb"
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "张三", "通过-name 告诉server你是谁")

func main() {
	// connect to server
	flag.Parse()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.Dial failed,err:", err)
		return
	}
	defer conn.Close()
	// create client
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// name := `张三`
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		fmt.Println("client.SayHello failed,err:", err)
	}
	fmt.Println("response:", resp.Reply)
}
