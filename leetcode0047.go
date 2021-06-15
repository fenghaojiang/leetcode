func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	visit := make([]bool, n)
	var backTrack func(curIndex int, arr []int)
	res := make([][]int, 0)
	backTrack = func(curIndex int, arr []int) {
		if curIndex == n {
			res = append(res, arr)
			return
		}
		for i, v := range nums {
			if visit[i] || i > 0 && !visit[i-1] && nums[i-1] == v {
				continue
			}
			arr = append(arr, v)
			visit[i] = true
			backTrack(curIndex+1, arr)
			visit[i] = false
			arr = append([]int{}, arr[:len(arr)-1]...)
		}
	}
	backTrack(0, []int{})
	return res
}