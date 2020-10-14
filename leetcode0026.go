package main

import "fmt"

//双指针
func removeDuplicates(nums []int) int {
	u := 0
	for i:= 1; i < len(nums); i++ {
		if nums[i] != nums[u] {
			u++
			nums[u] = nums[i]
		}
	}
	return u+1
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
}