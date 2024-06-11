package main

import "fmt"

func main() {
	ret := addBinary("000", "110")
	fmt.Println(ret)
}

func addBinary(a string, b string) string {
	// 尾数起每位计算
	len_a := len(a)
	len_b := len(b)

	var len int
	if len_a > len_b {
		len = len_a
	} else {
		len = len_b
	}

	i := len_a - 1
	j := len_b - 1
	carry := 0
	ret := make([]byte, len)
	for i >= 0 && j >= 0 {
		// 进行运算
		sum := int(a[i]) + int(a[j]) + carry
		fmt.Println(string(sum))
		ret[len-i-1] = byte(sum % 2)
		carry = sum / 2
		// i++ j++
		i--
		j--
	}
	fmt.Println(ret)
	// 处理不对位的数
	if len_a == len_b {
		// 只需处理进位
		if carry == 0 {
			return string(ret)
		} else {
			return string(ret)
		}
	}
	return string(ret)
}
