func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for _, r := range s {
		arr[r-'a']++
	}
	for _, r := range t {
		arr[r-'a']--
		if arr[r-'a'] < 0 {
			return false
		}
	}
	return true
}