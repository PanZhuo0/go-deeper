package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int //行,列
	var subboxes [3][3][9]int   //子盒子
	for i, row := range board {
		for j, c := range row {
			//如果这个位置的值为空值
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++            //第i行的下标为[index]的数个数++
			columns[j][index]++         //第j行的下标为[index]的数个数++
			subboxes[i/3][j/3][index]++ //第i个盒子++

			// 如果第i行中 index对应的字符出现了2次
			if rows[i][index] > 1 {
				return false
			}

			// 如果第j列中 index对应的字符出现了2次
			if columns[j][index] > 1 {
				return false
			}

			// 如果在本子盒子中中index出现过2次
			if subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return false
}
