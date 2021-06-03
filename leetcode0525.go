func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	var cnt, maxLen int
	for i, num := range nums {
		if num == 1 {
			cnt++
		} else {
			cnt--
		}
		if index, ok := mp[cnt]; ok {
			maxLen = max(maxLen, i-index)
		} else {
			mp[cnt] = i
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}