package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	n := 100000000

	// 改进的埃筛法
	eratothenes_improve(n)
	// 埃筛法
	eratothenes(n)
	// 改进的传统法
	traditiona_improve(n)
}
func eratothenes_improve(n int) {
	// bitmap
	notPrime := make([]bool, n)
	count := 0
	now := time.Now()
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		} else {
			// 没对应记录(是素数)
			count++
			for j := i * i; j < n; j += i { //i=2 j =2 j=4 j=6 j=8 这些都会被标记为合数
				notPrime[j] = true
			}
		}
	}
	duration := time.Since(now)
	fmt.Println(duration)
	fmt.Println(count)
}
func eratothenes(n int) {
	// bitmap
	notPrime := make([]bool, n)
	count := 0
	now := time.Now()
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		} else {
			// 没对应记录(是素数)
			count++
			for j := 2 * i; j < n; j += i { //i=2 j =2 j=4 j=6 j=8 这些都会被标记为合数
				notPrime[j] = true
			}
		}
	}
	duration := time.Since(now)
	fmt.Println(duration)
	fmt.Println(count)
}

func traditiona_improve(n int) {
	count := 0
	now := time.Now()
	for i := 2; i < n; i++ {
		isPrime := true
		// 使用小于平方相比全部判断 性能提高了约100倍
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			// for j := 2; j < i; j++ {
			if i%j == 0 {
				// 不是素数
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	dure := time.Since(now)
	fmt.Println(dure)
	fmt.Println(count)
}
