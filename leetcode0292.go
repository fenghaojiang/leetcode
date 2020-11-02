package main

import (
	"fmt"
)

func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	num := 4
	fmt.Println(canWinNim(num))
}
