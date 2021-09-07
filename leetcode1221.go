func balancedStringSplit(s string) int {
	var res, temp int
	for _, c := range s {
		if c == 'L' {
			temp++
		} else {
			temp--
		}
		if temp == 0 {
			res++
		}
	}
	return res
}