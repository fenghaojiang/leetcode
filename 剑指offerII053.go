func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}