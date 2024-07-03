package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sf/client/sf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
	服务流

客户端实现
*/
func main() {
	// 1. new grpc client
	conn, err := grpc.NewClient(`localhost:8080`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. pb client
	c := sf.NewGreeterClient(conn)
	// 3. call
	/* server端会响应流式的结果,直到server端主动关闭流 */
	stream, err := c.LotsOfReplies(context.Background(), &sf.Request{
		Name: "zhangsan",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	// client需要从steam中读取server端传过来的数据
	count := 1
	for {
		res, err := stream.Recv() //从流中获取server端的响应,每个响应都是.proto文件中定义的Response类型
		if err == io.EOF {
			// 读取结束
			break
		}

		if err != nil {
			// 从流中receive信息出错
			log.Fatal(err)
			return
		}
		fmt.Println("offset:", count, ` 从stream中获取到响应:`, res.GetResponse())
		count++
	}
}
