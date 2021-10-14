func peakIndexInMountainArray(arr []int) int {
	start, end := 0, len(arr)-2
	var res int
	for start <= end {
		mid := (start + end) >> 1
		if arr[mid] > arr[mid+1] {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}
