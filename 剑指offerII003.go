func countBits(n int) []int {
	res := []int{}
	for i := 0; i <= n; i++ {
		res = append(res, countOnes(i))
	}
	return res
}

func countOnes(num int) int {
	var ones int
	for num > 0 {
		ones++
		num &= num - 1
	}
	return ones
}