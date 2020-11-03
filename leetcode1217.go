package main

import (
	"fmt"
)

func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		return odd
	} else {
		return even
	}
}

func main() {
	chips := []int{2, 2, 2, 3, 3}
	fmt.Println(minCostToMoveChips(chips))
}
