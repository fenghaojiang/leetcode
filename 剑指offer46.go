func translateNum(num int) int {
	str := strconv.Itoa(num)
	dp := make([]int, len(str)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(str); i++ {
		if str[i-2:i] <= "25" && str[i-2:i] >= "10" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)]
}