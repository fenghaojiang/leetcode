func checkSubarraySum(nums []int, k int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	left := 0
	right := len(nums) - 1
	for right-left >= 2 {
		if sum%k == 0 {
			return true
		}
		if right == 0 {
			sum -= nums[left]
			left++
		} else {
			sum -= nums[right]
			right--
		}
	}
	return false //TODO
}

没有AC的版本回去再看看