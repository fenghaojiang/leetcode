package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	b := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(a, b))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i //把值的下标存放在map中
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY //都出现在map中
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
