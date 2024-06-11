package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// capacity = 2
	c := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	// sender
	go func() {
		defer wg.Done()

		c <- `foo`
		c <- `bar`
	}()

	// receiver
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Msg:", <-c)
		fmt.Println("Msg:", <-c)
	}()

	wg.Wait()
}
