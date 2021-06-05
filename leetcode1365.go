func smallerNumbersThanCurrent(nums []int) []int {
	times := [101]int{}
	for _, num := range nums {
		times[num]++
	}
	for i := 0; i < 100; i++ {
		times[i+1] += times[i]
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = times[v-1]
		}
	}
	return ans
}