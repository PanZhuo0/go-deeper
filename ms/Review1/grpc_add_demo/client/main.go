package main

import (
	"context"
	"demo/add/client/add"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/* client端 */
func main() {
	// 1. new grpc client connetion
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. new pb client
	c := add.NewAdderClient(conn)
	// 3. 调用方法
	reply, err := c.Add(context.Background(), &add.AddRequest{
		X: 10,
		Y: 20,
	})
	fmt.Println("send request to rpc server ...")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("response from server:", reply.Sum)
}
