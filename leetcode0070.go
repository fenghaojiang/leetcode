func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	stairNums := make([]int, n+1)
	stairNums[1] = 1
	stairNums[2] = 2
	for i := 3; i < n+1; i++ {
		stairNums[i] = stairNums[i-1] + stairNums[i-2]
	}
	return stairNums[n]
}