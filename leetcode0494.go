func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}