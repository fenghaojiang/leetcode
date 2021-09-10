func chalkReplacer(chalk []int, k int) int {
	l := len(chalk)
	if l == 0 {
		return 0
	}
	total := 0
	for _, v := range chalk {
		total += v
	}
	k = k % total
	i := 0
	for k > 0 {
		if k < chalk[i] {
			return i
		}
		k -= chalk[i]
		i++
		if i == l {
			i = 0
		}
	}
	return i
}