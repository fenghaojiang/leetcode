func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i1, i2 := 0, 0
    l1, l2 := len(nums1), len(nums2)
    res := make([]int, 0)
    for i1 < l1 && i2 < l2 {
        if nums1[i1] < nums2[i2] {
            i1++
        } else if nums2[i2] < nums1[i1] {
            i2++
        } else {
            res = append(res, nums1[i1])
            i1++
            i2++
        }
    }
    return res
}


