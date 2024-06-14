package main

import (
	"bufio"
	"context"
	"flow/server/flow"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flow_server()
	flow_client()
	bi_flow()
}

func bi_flow() {
	// client send Msg to server and receive response from server
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := flow.NewFlowRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute) //the stream will last 2 minutes
	defer cancel()
	// acquire stream object
	stream, err := c.Bi(ctx)
	if err != nil {
		panic(err)
	}
	// wait signal for stop stream
	waitc := make(chan struct{})

	// goroutine : receive Msg from server
	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				// server stop stream because EOF
				close(waitc)
				return
			}
			if err != nil {
				// receive Msg failed
				panic(err)
			}
			// receive succeed
			fmt.Println("AI:", reply.GetReturn())
		}
	}()

	// keeping send Msg from STDIN to server
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n') //read Msg from stdin by line
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// else
		if err := stream.Send(&flow.HelloRequest{Name: cmd}); err != nil {
			// send Msg to server failed
			panic(err)
		}
	}
	stream.CloseSend() //close stream
}

func flow_client() {
	// make a grpc connetion
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// grpc request
	c := flow.NewFlowRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Reqs(ctx)
	if err != nil {
		// get request flow failed
		panic(err)
	}
	names := []string{"zhangsan", "李四", "wangwu "}
	for _, name := range names {
		err := stream.Send(&flow.HelloRequest{Name: name})
		if err != nil {
			// send failed
			panic(err)
		}
	}
	// neet receive result
	res, err := stream.CloseAndRecv()
	if err != nil {
		// if got no response
		panic(err)
	}
	fmt.Println("client flow,cilent get result:", res.Return)
}

func flow_server() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := flow.NewFlowRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Replies(ctx, &flow.HelloRequest{Name: "name_zhangsan"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// stream result
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// 流结束
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println("Server flow:client get reply:", res.GetReturn())
	}

}
