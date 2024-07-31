package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// 1. New一个Dialer
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial(`ws://localhost:8000/ws`, nil) //这里Header可以实现用户发送特定的Header给Server
	if err != nil {
		log.Fatalln("Dial WS failed,err:", err)
	}
	defer conn.Close() //用户退出后，自动关闭连接
	// 2.发送Msg("你好")给WebSocket服务器端,并等待服务器端的响应
	conn.WriteMessage(websocket.TextMessage, []byte("你好!"))
	_, msgData, err := conn.ReadMessage()
	if err != nil {
		log.Fatalln(err)
	}
	// 控制台输出服务器端响应的结果
	fmt.Println(string(msgData))
}
