package main

import (
	"bi/client/bi"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
	双向流

server端实现
*/
type server struct {
	bi.UnimplementedGretterServer
}

func main() {
	// 1.端口设置
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. new grpc server
	s := grpc.NewServer()
	// 3. register server
	bi.RegisterGretterServer(s, &server{})
	// 4. 监听
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (s *server) BiTest(stream bi.Gretter_BiTestServer) error {
	// 1.持续从流中获取数据，一旦获取就响应数据到流中
	// 2.直到客户端主动关闭流，才结束
	// 3.出错如何处理? 客户端出错，服务端出错
	count := 1
	for {
		// 从stream中获取
		req, err := stream.Recv() //这里会阻塞 (通道 <- chan)
		if err != nil {
			// 如果接收失败，应该响应给client错误
			log.Fatal(err)
			return err
		}
		if err == io.EOF {
			// 如果客户端关闭了流,客户端只需响应nil,即可
			return nil
		}
		// 将结果响应到stream中
		stream.Send(&bi.Response{
			Reply: "offset:" + fmt.Sprint(count) + "  你好," + req.GetName(),
		})
		count++
	}
}
