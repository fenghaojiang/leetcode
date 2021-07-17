func maxSubArray(nums []int) int {
	var max = -math.MaxInt32
	temp := 0
	for _, v := range nums {
		if temp < 0 {
			temp = 0
		}
		temp += v
		if temp > max {
			max = temp
		}
	}
	return max
}