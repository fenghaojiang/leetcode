func groupAnagrams(strs []string) [][]string {
	charMap := make(map[[26]int][]string)
	for _, s := range strs {
		cnt := [26]int{}
		for _, r := range s {
			cnt[r-'a']++
		}
		charMap[cnt] = append(charMap[cnt], s)
	}
	result := make([][]string, 0)
	for _, v := range charMap {
		result = append(result, v)
	}
	return result
}