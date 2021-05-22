func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		si, sj := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return si+sj >= sj+si
	})
	res := ""
	zero := 0
	for _, num := range nums {
		if num == 0 {
			zero++
		}
		res += strconv.Itoa(num)
	}
	if zero == len(nums) {
		return "0"
	}
	return res
}