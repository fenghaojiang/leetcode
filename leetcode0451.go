func frequencySort(s string) string {
	freq := make(map[byte]int)
	maxfreq := 0
	for i := range s {
		freq[byte(s[i])]++
		maxfreq = max(maxfreq, freq[byte(s[i])])
	}
	buckets := make([][]byte, maxfreq+1)
	for ch, t := range freq {
		buckets[t] = append(buckets[t], ch)
	}
	res := make([]byte, 0, len(s))
	for i := maxfreq; i > 0; i-- {
		for _, c := range buckets[i] {
			res = append(res, bytes.Repeat([]byte{c}, i)...)
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}