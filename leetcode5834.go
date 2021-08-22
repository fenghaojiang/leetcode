func minTimeToType(word string) int {
	word = "a" + word
	res := 0
	for i := 1; i < len(word); i++ {
		res += getMinDest(word[i-1], word[i]) + 1
	}
	return res
}

func getMinDest(cur byte, dest byte) int {
	if cur > dest {
		if int(cur-dest) < 26-int(cur-dest) {
			return int(cur - dest)
		} else {
			return 26 - int(cur-dest)
		}
	} else {
		if int(dest-cur) < 26-int(dest-cur) {
			return int(dest - cur)
		} else {
			return 26 - int(dest-cur)
		}
	}
}