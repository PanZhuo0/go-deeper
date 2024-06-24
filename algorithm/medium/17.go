package main

import "fmt"

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

var combination []string //全局结果

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combination = []string{} //初始化
	helper(digits, 0, "")
	return combination
}

func helper(digits string, index int, c string) {
	if index == len(digits) {
		// index==len(digits)表示已经获取了所有结果
		combination = append(combination, c)
	} else {
		digit := digits[index]
		letters := m[digit] //第i个数字对应的字母集合的值
		for i := 0; i < len(letters); i++ {
			helper(digits, index+1, c+letters[i])
		}
	}
}

func main() {
	fmt.Println(letterCombinations(`22`))
}
