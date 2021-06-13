func makeEqual(words []string) bool {
	wordMap := make(map[byte]int)
	for _, word := range words {
		for _, c := range word {
			wordMap[byte(c)]++
		}
	}
	l := len(words)
	for _, value := range wordMap {
		if value%l != 0 {
			return false
		}
	}
	return true
}
