package main

import "fmt"

func main() {
	ret := addBinary("100", "110010")
	fmt.Println(ret)
}

func addBinary(a string, b string) string {
	len_a := len(a)
	len_b := len(b)
	i := len_a - 1
	j := len_b - 1

	carry := 0
	ret := []byte{}
	for i >= 0 && j >= 0 {
		// 尾数相加
		sum := int(a[i]-'0'+b[j]-'0') + carry
		// sum can be 0 / 1 / 2 / 3
		ret = append([]byte{byte(sum%2) + '0'}, ret...)
		carry = sum / 2
		i--
		j--
	}
	// 处理尾数
	// 3种情况 len_a == len_b len_a>len_b len_a <len_b
	if len_a == len_b {
		if carry > 0 {
			// 头插法
			ret = append([]byte{'1'}, ret...)
		}
		return string(ret)
	} else if len_a > len_b {
		i = len_a - len_b - 1
		for i >= 0 {
			sum := int(a[i]-'0') + carry
			// 头插法
			ret = append([]byte{byte(sum%2) + '0'}, ret...)
			carry = sum / 2
			i--
		}
		// 处理尾数
		if carry > 0 {
			ret = append([]byte{'1'}, ret...)
		}
		return string(ret)
	} else {
		j = len_b - len_a - 1 //2
		for j >= 0 {
			sum := int(b[j]-'0') + carry
			// 头插法
			ret = append([]byte{byte(sum%2) + '0'}, ret...)
			carry = sum / 2
			j--
		}
		if carry > 0 {
			ret = append([]byte{'1'}, ret...)
		}
	}

	return string(ret)
}
