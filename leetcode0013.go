func romanToInt(s string) int {
	var charMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	l := len(s)
	ans := 0
	for i, v := range s {
		if i < l-1 && charMap[byte(v)] < charMap[s[i+1]] {
			ans -= charMap[byte(v)]
		} else {
			ans += charMap[byte(v)]
		}
	}
	return ans
}