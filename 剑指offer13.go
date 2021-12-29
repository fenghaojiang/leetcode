func movingCount(m int, n int, k int) int {
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || visit[i][j] || (i%10+i/10+j%10+j/10) > k {
			return 0
		}
		visit[i][j] = true
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1) + 1
	}
	return dfs(0, 0)
}