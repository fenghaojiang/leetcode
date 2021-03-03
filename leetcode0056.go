package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	fmt.Println(merge(nums))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { //之前的最大值小于现在的最小值
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
