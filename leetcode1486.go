func xorOperation(n int, start int) int {
	temp := start
	for i := 1; i < n; i++ {
		start ^= (temp + 2*i)
	}
	return start
}