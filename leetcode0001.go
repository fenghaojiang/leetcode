package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		bias := target - nums[i]
		c, ok := m[bias]
		if ok {
			return []int{c, i}
		}
		m[nums[i]] = i
	}
	return []int{0,0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}