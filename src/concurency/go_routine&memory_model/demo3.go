package main

import "net/http"

func main() {
	go func() {
		done <- serverApp()
	}()
	for i := 0; i < cap(done); i++ {
		<-done
		close(stop)
	}
}

func serverApp(stop chan struct{}) error {
	// goroutine
	go func() {
		<-stop
		http.Shutdown()
	}()
	// call goroutine2
	return http.Listen()
}
