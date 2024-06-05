package main

func main() {
	go serverApp()
	select {}
}

func serverApp() {
	// 对于函数提供者，不可以假设自己被前台执行还是后台执行
	go func() {
		// http.listen
	}()
}
