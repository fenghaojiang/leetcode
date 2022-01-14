func printNumbers(n int) []int {
	s := int(math.Pow(10, float64(n)))
	res := []int{}
	for i := 1; i < s; i++ {
		res = append(res, i)
	}
	return res
}

//大数尼玛，暴力