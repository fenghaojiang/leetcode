func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	start, end := 0, len(people)-1
	var res int
	for start <= end {
		if people[start]+people[end] <= limit {
			start++
			end--
			res++
		} else {
			end--
			res++
		}
	}
	return res
}