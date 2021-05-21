func search(nums []int, target int) int {
	return binarySearch(0, len(nums)-1, target, nums)
}

func binarySearch(start, end, target int, nums []int) int {
	if start < 0 || end > len(nums)-1 {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target && target <= nums[end] {
		return binarySearch(mid+1, end, target, nums)
	} else if nums[mid] > target && target >= nums[start] {
		return binarySearch(start, mid-1, target, nums)
	}
	return -1
}