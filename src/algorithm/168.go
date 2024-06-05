package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
}

func convertToTitle(columnNumber int) string {
	ret := []byte{}
	for columnNumber > 0 {
		x := columnNumber % 26
		ret = append(ret, byte(x+'A'-1))
		columnNumber = columnNumber / 26
	}
	return string(ret)
}
