func minTimeToVisitAllPoints(points [][]int) int {
	var ans int
	x0, y0 := points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		ans += max(abs(x1-x0), abs(y1-y0))
		x0, y0 = x1, y1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}