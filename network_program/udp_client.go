package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// 1.连接到UDP 服务端
	conn, err := net.DialUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)}, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 20000}) //传输方式 本地地址 外部地址
	if err != nil {
		log.Fatal(err) //连接失败
	}
	defer conn.Close()
	// 2.写入数据到UDP
	sendData := []byte("hello udp")
	_, err = conn.Write(sendData)
	if err != nil {
		log.Fatalln(err) //写入数据到socket失败
	}
	// 3.从UDP中返回数据
	data := make([]byte, 1e3)
	n, addr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Fatalln(err) //读取服务端socket的返回失败
	}
	fmt.Printf("收到Addr:%v 的服务器的回复:%v 总量:%v bytes", addr, string(data), n)
}
