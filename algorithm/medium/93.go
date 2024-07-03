package main

import (
	"strconv"
	"strings"
)

func main() {

}

func restoreIpAddresses(s string) []string {
	path := make([]string, len(s))
	res := make([]string, 0)
	var dfs func(string, int)
	dfs = func(s string, startIdx int) {
		// 结束了
		if len(path) == 4 {
			// 分隔符在末尾
			if startIdx == len(s) {
				str := strings.Join(path, ".") //在对应分隔符位置加.
				res = append(res, str)
			}
			return
		}
		for i := startIdx; i < len(s); i++ {
			if i == startIdx && s[startIdx] == '0' {
				// 含有前导0,无效
				break
			}
			str := s[startIdx : i+1] //尝试分隔
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				// if 合法
				path = append(path, str)  //push
				dfs(s, i+1)               //下一次分隔
				path = path[:len(path)-1] //pop
			} else {
				// 不合法，分隔符再往后面放更不合法，因此break掉
				break
			}
		}
	}
	dfs(s, 0)
	return res
}
