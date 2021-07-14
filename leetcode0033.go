func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	start, end := 0, l-1
	for start <= end {
		mid := start + end
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] { //左边有序
			if nums[0] <= target && target < nums[mid] { //在左边
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { //右边有序
			if nums[mid] <= target && target < nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}