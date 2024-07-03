package main

import (
	"cf/server/cf"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
客户端流
client端实现
*/
func main() {
	// 1. new grpc client
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. new pb client
	s := cf.NewCFClient(conn)
	stream, err := s.Gretter(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 3. 将请求批量的写入到流中
	names := []string{`张三`, `李四`, `王五`}
	fmt.Println("client send msg to server by Stream")
	for _, name := range names {
		err := stream.Send(&cf.Request{Name: name})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 4.client端关闭流,并等待server端的响应 (我可以肯定底层用了通道)
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("client stream.CloseAndRecv", err)
		return
	}
	fmt.Println("client got response from server throught Stream:", resp)
	// 4.等待
}
