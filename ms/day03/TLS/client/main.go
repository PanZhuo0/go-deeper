package main

import (
	"context"
	"fmt"
	"time"
	"tls/client/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../server.crt", "")
	if err != nil {
		// load credential failed
		panic(err)
	}
	conn, err := grpc.NewClient("localhost:8080",
		grpc.WithTransportCredentials(creds),
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		// new gRPC client failed
		panic(err)
	}
	defer conn.Close()
	c := tls.NewTClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.THello(ctx, &tls.TReq{Name: "name_zhangsan"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.Reply)
}
