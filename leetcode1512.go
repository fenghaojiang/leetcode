package main

import "fmt"


func numIdenticalPairs(nums []int) int {
	var cnt int
	m := make(map[int]int)
	for _, num := range nums {
		if val, ok := m[num]; ok {
			cnt += val
		}
		m[num]++
	}
	return cnt 
}

func main() {
	nums := []int{1,2,3,1,1,3}
	fmt.Println(numIdenticalPairs(nums))
}