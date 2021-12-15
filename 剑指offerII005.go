func maxProduct(words []string) int {
	mask := make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			mask[i] |= 1 << (c - 'a')
		}
	}
	var res int
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if mask[i]&mask[j] == 0 {
				if len(words[i])*len(words[j]) > res {
					res = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return res
}