package main

import (
	"fmt"
	"sync"
)

type Config struct {
	a []int
}

func main() {
	cfg := &Config{}

	// 1 writer
	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()
	var wg sync.WaitGroup

	// 4 reader
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				fmt.Println(cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
