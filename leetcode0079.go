func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	if m*n < len(word) {
		return false
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		board[i][j] = '0'
		res := dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		board[i][j] = word[k]
		return res
	}
	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}