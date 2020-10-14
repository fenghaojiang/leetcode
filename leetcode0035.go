package main

import "fmt"

//二分
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (right + left) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1 
		}
	}
	return left
}

func main() {
	nums := []int{1,3,5,6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}