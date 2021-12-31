func isStraight(nums []int) bool {
	m := make(map[int]struct{})
	var max, min = 0, 14
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		if _, ok := m[nums[i]]; ok {
			return false
		}
		m[nums[i]] = struct{}{}
	}
	return max-min < 5
}
  


