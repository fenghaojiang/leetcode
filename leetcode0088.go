func merge(nums1 []int, m int, nums2 []int, n int) {
	h1, h2 := m-1, n-1
	cur := m + n - 1
	for h1 >= 0 && h2 >= 0 {
		if nums1[h1] > nums2[h2] {
			nums1[cur] = nums1[h1]
			h1--
		} else {
			nums1[cur] = nums2[h2]
			h2--
		}
		cur--
	}

	if h1 >= 0 {
		for ; cur >= 0 && h1 >= 0; cur-- {
			nums1[cur] = nums1[h1]
			h1--
		}
	}

	if h2 >= 0 {
		for ; cur >= 0 && h2 >= 0; cur-- {
			nums1[cur] = nums2[h2]
			h2--
		}
	}
}