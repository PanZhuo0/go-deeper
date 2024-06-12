package main

import (
	"client/homework"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.Dial failed,err:", err)
	}
	// defer close connection
	defer conn.Close()
	// new Client
	c := homework.NewHomeWorkClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Call Add(RPC)
	resp, err := c.Add(ctx, &homework.HomeWokrRequest{X: 99, Y: 999})
	if err != nil {
		fmt.Println("c.Add failed,err:", err)
	}
	fmt.Println("Answer:", resp.Sum)
}
