func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var backTrack func(start int, sum int, temp []int)
	backTrack = func(start int, sum int, temp []int) {
		if sum == 0 {
			res = append(res, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > sum {
				continue
			}
			temp = append(temp, candidates[i])
			backTrack(i+1, sum-candidates[i], temp)
			temp = append([]int{}, temp[:len(temp)-1]...)
		}
	}
	backTrack(0, target, []int{})
	return res
}