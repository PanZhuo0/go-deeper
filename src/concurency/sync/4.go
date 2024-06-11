package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)
	var mu sync.Mutex
	// go routine 1,it will try it's best to require mutex
	goroutine1 := 0
	goroutine2 := 0
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// difference: go routine 1 will mute firstly,then Sleep(100ms)
				// 				while go routine 2 will sleep 100ms,then mutex
				// go routine1 每次持100ms锁(不退避)
				// go routine2 每100ms持一次锁(持锁前，先退避100ms)
				goroutine1++
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// go routine2
	for i := 0; i < 10; i++ { //after 10 times go routine2 acquired mutex The program will end
		goroutine2++
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true

	fmt.Println(goroutine1)
	fmt.Println(goroutine2)
}
