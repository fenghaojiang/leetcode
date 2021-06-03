func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	mp := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if index, ok := mp[remainder]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}