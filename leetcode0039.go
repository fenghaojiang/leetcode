package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	comb := []int{}
	ans := make([][]int, 0)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
