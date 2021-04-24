package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	var cnt int
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})
	for _, v := range costs {
		if coins <= 0 {
			return cnt
		}
		if coins >= v {
			cnt++
			coins -= v
		}
	}
	return cnt
}
func main() {
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20))
}
