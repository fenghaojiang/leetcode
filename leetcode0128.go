func longestConsecutive(nums []int) int {
	visit := make(map[int]bool)
	for _, v := range nums {
		visit[v] = true
	}
	maxStep := 0
	for key := range visit {
		if !visit[key-1] {
			currentNum := key
			step := 1
			for visit[currentNum+1] {
				currentNum++
				step++
			}
			if maxStep < step {
				maxStep = step
			}
		}
	}
	return maxStep
}