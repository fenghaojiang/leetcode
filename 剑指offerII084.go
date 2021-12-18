func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	var backTrack func(temp []int)
	l := len(nums)
	visit := make([]bool, l)
	sort.Ints(nums)
	backTrack = func(temp []int) {
		if len(temp) == l {
			res = append(res, temp)
			return
		}
		for i := range nums {
			if visit[i] || (i > 0) && nums[i-1] == nums[i] && visit[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			visit[i] = true
			backTrack(temp)
			visit[i] = false
			//temp = temp[:len(temp)-1]
			temp = append([]int{}, temp[:len(temp)-1]...)
		}
	}
	backTrack([]int{})
	return res
}