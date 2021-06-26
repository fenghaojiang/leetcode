func canBeIncreasing(nums []int) bool {
	var nums1 = nums
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			nums2 = append(nums2[:i-1], nums2[i:]...)
			nums1 = append(nums1[:i], nums1[i+1:]...)
			break
		}
	}
	return isIncreasing(nums1) || isIncreasing(nums2)
}

func isIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}