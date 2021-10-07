func countSegments(s string) int {
	var cnt int
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			cnt++
		}
	}
	return cnt
}