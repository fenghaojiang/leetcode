func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	maxn, minn := -math.MaxInt64, math.MaxInt64
	right, left := -1, -1
	for i, v := range nums {
		if maxn > v {
			right = i
		} else {
			maxn = v
		}
		if minn < nums[l-i-1] {
			left = l - i - 1
		} else {
			minn = nums[l-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

