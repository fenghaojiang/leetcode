func lengthOfLongestSubstring(s string) int {
	//滑动窗口+hash
	if len(s) == 0 {
		return 0
	}
	res := 1
	set := make(map[byte]struct{})
	set[s[0]] = struct{}{}
	left := 0
	for right := 1; right < len(s); {
		if _, ok := set[s[right]]; !ok {
			set[s[right]] = struct{}{}
			right++
		} else {
			//往右滑动一个元素
			delete(set, s[left])
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}