func constructArr(a []int) []int {
	n := len(a)
	left := 1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left
		left *= a[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= a[i]
	}
	return res
}