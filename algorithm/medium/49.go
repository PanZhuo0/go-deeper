package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{`eat`, `tea`, `tan`, `ate`, `nat`, `bat`}))
}

func groupAnagrams(strs []string) [][]string {
	/*
		key: [26]int  //每个字符出现次数的记录数组 ,Go 居然可以把数组当作key,厉害
		value: string字符串数组
	*/
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++ //计数++
		}
		m[cnt] = append(m[cnt], str)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
