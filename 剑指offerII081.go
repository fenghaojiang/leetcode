func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	nums := make([]int, 0)
	sort.Ints(candidates)
	var backTrack func(sum int, index int)
	backTrack = func(sum int, index int) {
		if sum == target {
			res = append(res, nums)
			return
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			nums = append(nums, candidates[i])
			backTrack(sum+candidates[i], i)
			nums = append([]int{}, nums[:len(nums)-1]...)
		}
	}
	backTrack(0, 0)
	return res
}