package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close() //关闭连接
	for {
		var data [1024]byte
		n, addr, err := l.ReadFromUDP(data[:])
		if err != nil {
			log.Fatalln(err)
		}
		// UDP没有建立连接，所有收发数据需要确认你的接收方的IP Addr
		log.Printf("接收到UDP Addr:%v ,数据量:%v bytes, 数据:%s", addr, n, string(data[:n]))
		_, err = l.WriteToUDP(data[:n], addr) //数据+地址
		if err != nil {
			log.Fatalln(err)
		}
	}
}
