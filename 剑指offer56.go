func singleNumbers(nums []int) []int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	div := 1
	// 找到全员异或结果中为1的位。
	// 即可将全员分成该位为1和该位为0的两组
	// 出现过两次的数字在同一组，出现过一次的在不同组
	for (div & res) == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		// 通过与运算进行分组
		if (v & div) == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}