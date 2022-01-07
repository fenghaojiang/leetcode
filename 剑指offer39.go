func majorityElement(nums []int) int {
	var max int
	var cnt int
	for i := range nums {
		if cnt == 0 {
			max = nums[i]
		}
		if max == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return max
}  


