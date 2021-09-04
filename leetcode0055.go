func canJump(nums []int) bool {
	var maxLength int
	n := len(nums)
	for i := 0; i < n; i++ {
		if i <= maxLength {
			maxLength = max(maxLength, i+nums[i])
			if maxLength >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}