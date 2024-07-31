package main

import (
	"fmt"
	"log"
	"net/http"
)

// 一个HTTP服务
func main() {
	// 带证书认证的也就是HTTPS
	// http.ListenAndServeTLS()
	http.HandleFunc("/go", handler) //这里的函数会注册到默认的多路复用服务器中(DefualtServerMux)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("127.0.0.1:8000", nil) //提供HTTP服务的IP,以及端口
	if err != nil {
		log.Fatalln(err)
	}
	// http 域名其实就是IP+端口,域名的默认端口就是80端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取信息
	// log.Println(r)
	log.Printf(`proto:%v, url:%s ,method:%s`, r.Proto, r.URL, r.Method)
	log.Printf(`path:%v`, r.URL)
	// 回应数据
	w.Write([]byte("[Testing]---> 码神之路 Go"))
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	info := fmt.Sprintf("username:%s password:%s logined!", username, password)
	w.Write([]byte(info))
}
