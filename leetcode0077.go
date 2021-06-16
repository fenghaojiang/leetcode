func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backTrack func(curIndex int, nums []int)
	backTrack = func(curIndex int, nums []int) {
		if len(nums)+(n-curIndex+1) < k {
			return
		}
		if len(nums) == k {
			res = append(res, nums)
			return
		}
		for i := curIndex; i <= n; i++ {
			nums = append(nums, i)
			backTrack(i+1, nums)
			nums = append([]int{}, nums[:len(nums)-1]...)
		}
	}
	backTrack(1, []int{})
	return res
}