func search(nums []int, target int) int {
	var res int
	var binarySearch func(left, right int)
	binarySearch = func(left, right int) {
		if left > right {
			return
		}
		mid := (left + right) >> 1
		if nums[mid] == target {
			res++
			binarySearch(left, mid-1)
			binarySearch(mid+1, right)
		} else if nums[mid] > target {
			binarySearch(left, mid-1)
		} else {
			binarySearch(mid+1, right)
		}
	}
	binarySearch(0, len(nums)-1)
	return res
}