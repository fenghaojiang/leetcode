func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(start int, subNums []int)
	dfs = func(start int, subNums []int) {
		res = append(res, subNums)
		for i := start; i < len(nums); i++ {
			subNums = append(subNums, nums[i])
			dfs(i+1, subNums)
			subNums = append([]int{}, subNums[:len(subNums)-1]...)
		}
	}
	dfs(0, []int{})
	return res
}