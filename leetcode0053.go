package main

import "fmt"

func main() {
	nums := []int{-1}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	temp := 0
	max := -99999999
	for i := 0; i < len(nums); i++ {
		if temp < 0 {
			temp = 0
		}
		temp += nums[i]
		if temp > max {
			max = temp
		}
	}
	return max
}
