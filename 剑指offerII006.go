func twoSum(numbers []int, target int) []int {
	var binarySearch func(left, right, target int) int
	binarySearch = func(left, right, target int) int {
		for left <= right {
			mid := (left + right) >> 1
			if numbers[mid] == target {
				return mid
			}
			if numbers[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	for i := range numbers {
		x := numbers[i]
		j := binarySearch(i+1, len(numbers)-1, target-x)
		if j != -1 {
			return []int{i, j}
		}
	}
	return []int{}
}