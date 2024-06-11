package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(2147483647))
}

func convertToTitle(columnNumber int) string {
	ret := []byte{}
	for columnNumber > 0 {
		// 这里的字符是1-index的，即以1开头
		// 因此在计算mod时，应该让数据-1，让整体变为0-index
		// 才能得到正确的mod
		x := (columnNumber - 1) % 26
		ret = append([]byte{byte(x + 'A')}, ret...)
		columnNumber = (columnNumber - 1) / 26
	}

	return string(ret)
}
