package main

import "fmt"

func main() {
	fmt.Println(grayCode(25))
}

/* 格雷编码 */
/* 归纳法 */
/* 格雷编码存在的规律
n:位数
当n=0,格雷编码序列:[0]
当n=x,格雷编码序列==G(n-1)<<1+G(n-1)倒序<<1+1


eg:
	n=2
		0 0
		0 1
		1 1
		1 0

	n=3
		keep same and left shift once
		0 0 0
		0 1 0
		1 1 0
		1 0 0
		----- add one for each
		0 0 1
		0 1 1
		1 1 1
		1 0 1

*/

func grayCode(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0)
	for n > 0 {
		n--
		for i := len(ans) - 1; i >= 0; i-- {
			ans[i] <<= 1                //先左移一位(上半部分)
			ans = append(ans, ans[i]+1) //再加一(下半部分,上半部分再加一)
		}
	}
	return ans
}

/* 公式法:跳过 */
