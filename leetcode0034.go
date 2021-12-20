func searchRange(nums []int, target int) []int {
	var res = []int{math.MaxInt32, math.MinInt32}
	var binarySearch func(left, right int)
	binarySearch = func(left, right int) {
		if left > right {
			return
		}
		mid := (left + right) >> 1
		if nums[mid] == target {
			if mid < res[0] {
				res[0] = mid
			}
			if mid > res[1] {
				res[1] = mid
			}
			binarySearch(left, mid-1)
			binarySearch(mid+1, right)
		} else if nums[mid] > target {
			binarySearch(left, mid-1)
		} else {
			binarySearch(mid+1, right)
		}
	}
	binarySearch(0, len(nums)-1)
	if res[0] == math.MaxInt32 && res[1] == math.MinInt32 {
		return []int{-1, -1}
	}
	return res
}