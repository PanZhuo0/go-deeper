package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// 1.启动一个HTTP服务
	http.HandleFunc("/ws", handle)
	http.ListenAndServe(":8000", nil)
}

var up = websocket.Upgrader{
	HandshakeTimeout: time.Second * 3, //握手超时时间
	ReadBufferSize:   1e3,             //读Buffer容量 1KB 1e3==>10**3
	WriteBufferSize:  1e3,             //写Buffer容量 1KB
}

func handle(w http.ResponseWriter, r *http.Request) {
	// 2.升级为websocket
	conn, err := up.Upgrade(w, r, nil) //第3个参数，可以让我们实现自定义协议的请求头
	if err != nil {
		log.Fatalln("升级websocket失败,err:", err)
	}
	_, msgByte, err := conn.ReadMessage()
	if err != nil {
		log.Println("conn.ReadMessage() FAILED,err:", err)
	}
	// 3.对数据进行处理
	str := string(msgByte)
	str = strings.Trim(str, "\r\n")
	conn.WriteMessage(websocket.TextMessage, []byte(str))
}
