func hammingWeight(num uint32) int {
	var flag uint32 = 1
	var cnt int
	for i := 0; i < 32; i++ {
		if (flag<<i)&num > 0 {
			cnt++
		}
	}
	return cnt
}