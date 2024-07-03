package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(uniquePaths(5, 2))
}

func uniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
