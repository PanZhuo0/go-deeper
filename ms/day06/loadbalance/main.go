package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/* Client demo 
	使用gRPC内置的LB算法:round_robin 轮询
*/
func main() {
	conn, err := grpc.NewClient(`localhost:8080`,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
}
