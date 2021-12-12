func subArrayRanges(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	var backTrack func(index int)
	var res int64
	backTrack = func(index int) {
		if index == len(nums)-1 {
			return
		}
		var min = nums[index]
		var max = nums[index]
		for i := index; i < len(nums); i++ {
			if max < nums[i] {
				max = nums[i]
			}
			if min > nums[i] {
				min = nums[i]
			}
			res += int64(max - min)
		}
		backTrack(index + 1)
	}
	backTrack(0)
	return res
}