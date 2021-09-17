func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(index int) int {
		if index <= -1 || index >= n {
			return math.MinInt64
		}
		return nums[index]
	}

	start, end := 0, n-1
	for {
		mid := (start + end) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid-1) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
}