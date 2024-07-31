package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	// 1.连接websocket
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial(`ws://localhost:8080/ws`, nil)
	if err != nil {
		// 连接不上，则需要结束该程序
		log.Fatalln(err)
	}
	defer conn.Close() //defer 关闭websocket连接

	// 3.开启一个协程，用户不断读取服务器推送的消息
	go func() {
		for {
			_, msgData, err := conn.ReadMessage()
			if err != nil {
				log.Println("从服务器读取信息failed,err:", err.Error())
			}
			fmt.Println(string(msgData))
		}
	}()

	// 2.从标准输入获取输入,并将其发往websocket
	rd := bufio.NewReader(os.Stdin)
	for {
		// 2.1不断从标准输入获取信息
		str, err := rd.ReadString('\n')
		if err != nil {
			// 如果获取输入信息失败,记录日志,继续尝试获取
			log.Println("rd.ReadString() failed,err:", err.Error())
		}
		str = strings.Trim(str, "\r\n")
		// 2.2如果用户输入Q/q 则退出
		if strings.ToUpper(str) == `Q` {
			/*
				Issue:
				break/return都一样 不知道为什么，会在该用户退出前往服务器发送很多空消息
				conn.Close() //尝试显式的关闭conn,避免重复发送,也不行，大概是服务器端的问题
				目前来看，应该是server端会在websocket连接关闭后，仍然尝试从客户端读取数据,而且会读取成功(结果为空),直到conn实际上关闭(被GC)
			*/
			break
		}
		// 2.3如果用户输入的是普通信息,将数据写入到服务器
		err = conn.WriteMessage(websocket.TextMessage, []byte(str))
		if err != nil {
			// 如果将用户写入服务器失败,记录日志
			log.Println("conn.WriteMessage failed,err:", err.Error())
		}
	}

}
