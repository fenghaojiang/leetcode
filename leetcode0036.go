func isValidSudoku(board [][]byte) bool {
	showInRow := [9][9]int{}
	showInCol := [9][9]int{}
	showInBox := [3][3][9]int{}
	for i := range board {
		for j, c := range board[i] {
			if c == '.' {
				continue
			}
			index := c - '1'
			showInRow[i][int(index)]++
			showInCol[j][int(index)]++
			showInBox[i/3][j/3][int(index)]++
			if showInRow[i][int(index)] > 1 || showInCol[j][int(index)] > 1 || showInBox[i/3][j/3][int(index)] > 1 {
				return false
			}
		}
	}
	return true
}