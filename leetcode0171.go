func titleToNumber(columnTitle string) int {
	res := 0
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		res += int(k) * multiple
		multiple *= 26
	}
	return res
}