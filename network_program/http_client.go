package main

import (
	"log"
	"net/http"
)

func main() {
	// 可以不编写客户端、浏览器实际上可以当成一个HTTP的客户端
	resp, _ := http.Get(`http://localhost:8000/go`)
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		// 一直读取，直到读取完全部、EOF
		n, err := resp.Body.Read(buf)
		if err != nil {
			log.Printf("Response:%v ,Summary: %v bytes", string(buf), n)
			log.Fatalln(err)
		}
		log.Printf("Response:%v ,Summary: %v bytes", string(buf), n)
	}
}
