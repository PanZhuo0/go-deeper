package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	var ret = []string{}
	helper(0, 0, n, "", &ret)
	return ret
}

func helper(left, right, n int, r string, ret *[]string) {
	// 如果left==n right==n,所有括号用完了
	if left == n && right == n {
		*ret = append(*ret, r)
		return
	}
	// 如果left<n 增加一个左括号
	if left < n {
		helper(left+1, right, n, r+"(", ret)
	}
	if left > right && right < n {
		helper(left, right+1, n, r+")", ret)
	}
}
