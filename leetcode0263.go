package main

import (
	"fmt"
)

func isUgly(num int) bool {
	for i := 2; i < 6 && num > 0; i++ {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}

func main() {
	fmt.Println(isUgly(8))
}
