func toLowerCase(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			b[i] = s[i] + 32
		}
	}
	return string(b)
}