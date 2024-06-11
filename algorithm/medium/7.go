package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(1534236469))
}
func reverse(x int) int {
	// 变为字符串然后反转即可
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	bytex := []byte(strconv.Itoa(x))
	// strX反转
	for i := 0; i < len(bytex)/2; i++ {
		// exchange
		bytex[i], bytex[len(bytex)-1-i] = bytex[len(bytex)-1-i], bytex[i]
	}

	x64, err := strconv.ParseInt(string(bytex), 10, 32)
	if err != nil {
		// 超出表示范围
		return 0
	}
	x = int(x64)

	if negative {
		return -x
	}
	return x
}
