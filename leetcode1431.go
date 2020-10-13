package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := -1
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	res := make([]bool, len(candies))
	for i, v := range candies {
		if v + extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {
	candies := []int{2,3,5,1,3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}