package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", `0.0.0.0:20000`) //接收本机网卡下所有可用IP的请求,所有广播信号都接受
	// :20000 代表默认0.0.0.0:20000
	// 为了安全最好还是绑定一个固定的IP地址
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	defer conn.Close()
	for {
		rd := bufio.NewReader(conn)
		buf := make([]byte, 512)
		n, err := rd.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		str := string(buf[:n])
		fmt.Println("从客户端获取到的输入:", str)
		conn.Write(buf) //写回
	}
}
