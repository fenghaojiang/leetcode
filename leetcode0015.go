func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for first := range nums {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}