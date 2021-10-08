func findRepeatedDnaSequences(s string) []string {
	const L int = 10
	m := make(map[string]int)
	for i := 0; i <= len(s)-L; i++ {
		m[s[i:i+L]]++
	}
	res := make([]string, 0)
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}