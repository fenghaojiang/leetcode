func triangleNumber(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	var res int
	for i, v := range nums {
		k := i
		for j := i + 1; j < l; j++ {
			for k+1 < l && nums[k+1] < v+nums[j] {
				k++
			}
			res += max(k-j, 0)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}