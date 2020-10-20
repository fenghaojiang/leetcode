package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	res := make([]int,0)
	var max int
	var startIndex = 0
	var endIndex = k
	for endIndex <= len(nums)  {
		max = -999999
		for _, v := range nums[startIndex:endIndex] {
			if v > max {
				max = v
			}
		}
		res = append(res, max)
		startIndex++
		endIndex++
	} 
	return res
}

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}