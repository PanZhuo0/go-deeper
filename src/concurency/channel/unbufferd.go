package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		// sender sent a sting
		c <- `foo`
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		// receiver receive string
		fmt.Println("Msg:", <-c)
	}()

	wg.Wait()
}
