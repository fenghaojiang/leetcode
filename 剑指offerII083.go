func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backTrack func(temp []int)
	l := len(nums)
	visit := make([]bool, l)
	backTrack = func(temp []int) {
		if len(temp) == l {
			res = append(res, temp)
			return
		}
		for i := range nums {
			if visit[i] {
				continue
			}
			visit[i] = true
			temp = append(temp, nums[i])
			backTrack(temp)
			visit[i] = false
			temp = append([]int{}, temp[:len(temp)-1]...)
		}
	}
	backTrack([]int{})
	return res
}