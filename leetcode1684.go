func solve(s string) int {
	n := 0
	for _, v := range s {
		n |= 1 << (v - 'a') //"ac" -> 5   (000...101)
	}
	return n
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	ia := solve(allowed)
	for _, word := range words {
		iw := solve(word)
		if iw&ia == iw {
			count++
		}
	}
	return count
}