package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var A int = 0

func main() {
	// create 2 goroutine
	for routine := 1; routine <= 2; routine++ {
		wg.Add(1)
		go Routine(routine)
	}
	wg.Wait()
	fmt.Println("final count:", A)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := A
		time.Sleep(time.Nanosecond)
		value++
		A = value
	}
	wg.Done()
}
