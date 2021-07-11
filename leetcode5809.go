func countPalindromicSubsequence(s string) int {
	n := len(s)
	var res int
	for ch := 'a'; ch <= 'z'; ch++ {
		l, r := 0, n-1
		// 第一次出现的下标
		for l < n && byte(ch) != s[l] {
			l++
		}
		for r >= 0 && byte(ch) != s[r] {
			r--
		}
		if r-l < 2 {
			continue
		}
		charset := make(map[byte]struct{})
		for k := l + 1; k < r; k++ {
			charset[s[k]] = struct{}{}
		}
		res += len(charset)
	}
	return res
}