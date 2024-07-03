package main

import "fmt"

func main() {
	t := make([][]byte, 3)
	t[0] = []byte{'A', 'B', 'C', 'E'}
	t[1] = []byte{'S', 'F', 'C', 'S'}
	t[2] = []byte{'A', 'D', 'E', 'E'}
	fmt.Println(exist(t, "ABCCED"))
}

/* [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]] */
func exist(board [][]byte, word string) bool {
	// 1.initialize visited array
	vis := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	// 1.2 initialize a direction vector
	type pair struct{ x, y int }
	var dir = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 2.dfs function
	var dfs func(i, j, start int) bool //must declare first,so that we can use it in dfs func(recursively use)
	dfs = func(i, j, start int) bool {
		// 2.1 not avalid
		if board[i][j] != word[start] {
			return false
		}
		// 2.2 avalid,check passed
		if start == len(word)-1 {
			return true
		}

		// 2.3 need to check next element
		vis[i][j] = true
		defer func() { vis[i][j] = false }()
		for _, d := range dir {
			if newi, newj := i+d.x, j+d.y; 0 <= newi && newi < len(board) &&
				0 <= newj && newj < len(board[0]) && !vis[newi][newj] {
				// try next
				if dfs(newi, newj, start+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
