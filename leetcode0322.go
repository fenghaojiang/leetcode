package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	menoy := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		menoy[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && menoy[i] > menoy[i-c] + 1 {
				menoy[i] = menoy[i-c] + 1
			}
		}
	}
	if menoy[amount] == math.MaxInt32 {
		return -1
	}
	return menoy[amount]
}

func main() {
	coins := []int{1,2,5}
	fmt.Println(coinChange(coins, 11))
}
