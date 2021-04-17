package main

import (
	"fmt"
)

func maxDepth(s string) int {
	var cnt, max  int
	for _, r := range s {
		if r == '(' {
			cnt++
			if cnt > max {
				max = cnt 
			}
		} else if r == ')' {
			cnt--
		}
	}
	return max 
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}
