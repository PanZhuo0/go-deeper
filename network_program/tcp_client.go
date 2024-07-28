package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
	关于地址:

127.0.0.1:	本地环回地址
0.0.0.0:	广播地址,只要是该服务器网卡中的IP都可以接收(处于该网卡的广播域中)
*/
func main() {
	conn, err := net.DialTimeout("tcp", `192.168.1.70:20000`, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(os.Stdin)
	for {
		// 1.从标准输入获取数据
		str, err := rd.ReadString('\n')
		str = strings.Trim(str, "\r\n") //Attentino:(New Knowleage)-> 需要去除掉末尾的\r\n,才能得到输入的正确结果,\r\n的含义为回车return+换行
		if err != nil {
			log.Fatalln(err)
		}
		if strings.ToUpper(str) == "Q" {
			//输入为Q 或者q 则退出
			break
		}
		// 2.将数据输入
		conn.Write([]byte(str))
		buffer := make([]byte, 512)
		// 从socket连接中读取
		_, err = conn.Read(buffer)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("从服务端接收到响应response:", string(buffer))
	}
}
