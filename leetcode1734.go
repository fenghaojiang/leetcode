func decode(encoded []int) []int {
	first := 1
	n := len(encoded) + 1
	for i := 2; i <= n; i++ {
		first ^= i
	}
	for i := 1; i < n-1; i += 2 {
		first ^= encoded[i]
	}
	ans := make([]int, n)
	ans[0] = first
	for i := 0; i < n-1; i++ {
		ans[i+1] = ans[i] ^ encoded[i]
	}
	return ans
}