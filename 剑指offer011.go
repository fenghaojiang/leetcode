// 解题思路
// 本质上是找无序的地方
// nums[mid] > nums[right]表明在mid-right这一段开始乱序

// 代码

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := (left + right) >> 1
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}