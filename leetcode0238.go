func productExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}
	r := 1
	for i := length - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}
	return ans
}