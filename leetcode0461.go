package main

import "fmt"

func hammingDistance(x int, y int) int {
	result := x ^ y
	return countOne(result)
}

func countOne(num int) int {
	var cnt int 
	start := 0
	for start <31 {
		if num&(1<<start) > 0 {
			cnt++
		}
		start++
	}
	return cnt
}

func main() {
	fmt.Println(1<<1)
	fmt.Println(hammingDistance(1,4))
}