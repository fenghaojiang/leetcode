func maxValue(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	values := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		values[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			values[i][j] += max(values[i-1][j], values[i][j-1]) + grid[i-1][j-1]
		}
	}
	return values[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}