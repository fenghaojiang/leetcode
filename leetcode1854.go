func maximumPopulation(logs [][]int) int {
	nums := [101]int{}
	for _, life := range logs {
		for i := life[0]; i < life[1]; i++ {
			nums[i-1950]++
		}
	}
	var max, index int
	for i, year := range nums {
		if year > max {
			max = year
			index = i
		}
	}
	return index + 1950
}