func canBeIncreasing(nums []int) bool {
	var nums1 []int
	var nums2 []int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			nums2 = append(nums[:i-1], nums[i:]...)
			nums1 = append(nums[:i], nums[i+1:]...)
		}
	}
	return isIncreasing(nums1) || isIncreasing(nums2)
}

func isIncreasing(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return true
	}
	for i := 1; i < l; i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}