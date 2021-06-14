func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backTrack func(nums []int, target int)
	backTrack = func(nums []int, target int) {
		if len(nums) == k && target == 0 {
			res = append(res, nums)
			return
		}
		for i := 1; i < 10; i++ {
			if target >= i {
				nums = append(nums, i)
				backTrack(nums, target-i)
				nums = append([]int{}, nums[:len(nums)-1]...)
			}
		}
	}
	backTrack([]int{}, n)
	return res
}