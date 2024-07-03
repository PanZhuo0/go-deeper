package main

func main() {

}

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	ans := [][]string{}
	// 将整个n*n的矩阵赋值未true
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	// 倒序
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 这是在干嘛？ DP??
			f[i][j] = (s[i] == s[j] && f[i+1][j-1])
		}
	}
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 如果分割到了最后一个元素、（也就是回溯树的叶子节点）
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}
