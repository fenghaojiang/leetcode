package main

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	max := -1
	var index int
	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
