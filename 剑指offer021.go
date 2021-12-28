func exchange(nums []int) []int {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		if left < right && nums[left]&1 == 1 {
			left++
		}
		if left < right && nums[right]&1 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}