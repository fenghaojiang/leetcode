func leastMinutes(n int) int {
	speed := 1
	ans := 1
	for speed < n {
		speed *= 2
		ans++
	}
	return ans
}