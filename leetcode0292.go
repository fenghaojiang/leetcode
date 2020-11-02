package main

import (
	"fmt"
)

func findComplement(num int) int {
	temp := num
	n := 1
	for temp > 0 {
		n = n << 1
		temp = temp >> 1
	}
	n = n - 1
	return num ^ n
}

func main() {
	num := 5
	fmt.Println(findComplement(num))
}
