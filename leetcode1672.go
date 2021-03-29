package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	var max int = 0
	for i := 0; i < len(accounts); i++ {
		var temp int = 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func main() {
	a := [][]int{[]int{1,2,3}, []int{3,2,1}}
	fmt.Println(maximumWealth(a))
}
