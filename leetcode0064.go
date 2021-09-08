func minPathSum(grid [][]int) int {
	var infect func(i, j int)
	infect = func(i, j int) {
		if i == 0 && j == 0 {
			return
		}
		if i > 0 && j > 0 {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		} else if i > 0 {
			grid[i][j] += grid[i-1][j]
		} else if j > 0 {
			grid[i][j] += grid[i][j-1]
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			infect(i, j)
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}