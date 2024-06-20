package main

import "fmt"

func main() {
	fmt.Println(convert("AB", 1))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// Z字形变换
	flag := -1
	i := 0
	ss := make([]string, numRows)
	for k, _ := range s {
		ss[i] = ss[i] + string(s[k])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += ss[i]
	}
	return ret
}
