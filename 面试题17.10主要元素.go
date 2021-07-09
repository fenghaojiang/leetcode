func majorityElement(nums []int) int {
	var mainCnt int = 0
	var element int = -1
	l := len(nums) / 2
	for _, v := range nums {
		if mainCnt == 0 {
			element = v
		}
		if v != element {
			mainCnt--
		} else {
			mainCnt++
		}
		if mainCnt > l {
			return element
		}
	}
	mainCnt = 0
	for _, v := range nums {
		if v == element {
			mainCnt++
		}
	}
	if mainCnt*2 > len(nums) {
		return element
	}
	return -1
}