package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(array []int)
	hash := make(map[int]bool)
	dfs = func(array []int) {
		if len(array) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, array)
			res = append(res, temp)
			return 
		}
		for _, v := range nums {
			if !hash[v] {
				array = append(array, v)
				hash[v] = true
				dfs(array)
				hash[v] = false
				array = array[:len(array)-1]
			}
		}
	}
	array := make([]int, 0)
	dfs(array)
	return res
}



func main() {
	fmt.Println(permute([]int{1,2,3}))
}
