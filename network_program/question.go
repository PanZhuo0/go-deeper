package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1.怎么从stdin读取数据
	rd := bufio.NewReader(os.Stdin)
	str, err := rd.ReadString('\n')
	str = strings.Trim(str, "\r\n") //这里需要去除\r\n 因为ReadString()读取返回的数据总会在最后有\r\n
	if err != nil {
		panic(err)
	}
	fmt.Println("str:", str)
	// 2.net.Listent的地址
	/*
		0.0.0.0:20000 表示监听该网卡下的所有可用的IP的请求(广播帧)(默认为0.0.0.0)
			建议使用某个特定的地址,更安全
		127.0.0.1:20000 表示监听本机IP
		xxx.xxx.xxx.xxx :20000 表示监听本机某个网卡IP地址
	*/
}
