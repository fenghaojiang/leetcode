func checkIfPangram(sentence string) bool {
	mp := make(map[byte]struct{})
	for _, r := range sentence {
		mp[byte(r)] = struct{}{}
	}
	return len(mp) == 26
}