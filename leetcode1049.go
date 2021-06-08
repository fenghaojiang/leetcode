func lastStoneWeightII(stones []int) int {
	var sum int
	for _, v := range stones {
		sum += v
	}
	m := sum / 2
	dp := make([]bool, m+1)
	dp[0] = true
	for _, v := range stones {
		for j := m; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	for j := m; ; j-- {
		if dp[j] {
			return sum - 2*j
		}
	}
}