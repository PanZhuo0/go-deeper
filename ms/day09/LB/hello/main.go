package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(`consul://localhost:8500/hello?`,
		grpc.WithDefaultServiceConfig(`{"loadbalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(`get grpc Client failed,err:`, err)
		return
	}
	defer conn.Close()
}
