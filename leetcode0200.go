func numIslands(grid [][]byte) int {
	var res int
	var infect func(i, j int)
	infect = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		infect(i+1, j)
		infect(i-1, j)
		infect(i, j+1)
		infect(i, j-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				infect(i, j)
				res++
			}
		}
	}
	return res
}