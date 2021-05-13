package main

func numWays(steps int, arrLen int) int {
	dp := [502]int{1}
	for i := 1; i <= steps; i++ {
		now := 0
		for j := 0; j <= min(i, arrLen-1); j++ {
			now, dp[j] = dp[j], now
			dp[j] = ((now+dp[j])%1000000007 + dp[j+1]) % 1000000007
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
