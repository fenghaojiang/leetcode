package main

import "fmt"

//自己添加空集,之后每加一个集合形成新的集合添加到结果中去
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})

	for i := 0; i < len(nums); i++ {
		temp := make([][]int, 0)
		for _, arr := range result {
			ele := make([]int, 0)
			ele = append(ele, arr...)
			ele = append(ele, nums[i])
			temp = append(temp, ele)
		}
		for _, arr := range temp {
			result = append(result, arr)
		}
	}
	return result
}

func main() {

	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}