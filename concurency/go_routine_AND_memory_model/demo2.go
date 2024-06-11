package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		// done <- serveDebug(stop)
	}()
	go func() {
		// done <- serveApp(stop)
		// done<-serve()
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		// 任何一个serve 出错
		if err := <-done; err != nil {
			fmt.Println("error:%+v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop //wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
