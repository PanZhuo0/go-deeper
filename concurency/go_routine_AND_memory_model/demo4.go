package main

import (
	"fmt"
	"time"
)

func main() {
	search("123")
}

// 用户一般不能容忍过大的等待时间，因此需要设计超时机制
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

func process(term string) error {
	record, err := search(term)
	if err != nil {
		return err
	}
	fmt.Println("Received:", record)
	return nil
}
