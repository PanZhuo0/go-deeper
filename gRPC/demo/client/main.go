package main

import (
	"context"
	"fmt"

	pb "gRPC/demo/server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// build connection
	client := pb.NewSayHelloClient(conn)
	// call rpc
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "zhangsan"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.GetResponseMsg())
}
