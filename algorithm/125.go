package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func isPalindrome(s string) bool {
	// ASCII码中只取'0'-'9' 以及'a' ->'z' 'A'-'Z'
	// 替换字符
	var rep []byte
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			//保留
			rep = append(rep, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			// 大写字母变小写字母
			rep = append(rep, s[i]+32)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			// 保留
			rep = append(rep, s[i])
		} else {
			// 去除
		}
	}
	for i := 0; i < len(rep); i++ {
		if rep[i] != rep[len(rep)-i-1] {
			return false
		}
	}
	fmt.Println(len(rep))
	return true
}
