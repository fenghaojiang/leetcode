package main

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		if num&1 == 1{
			cnt++
		}
		num = num >> 1
	}
	return cnt
}

func main() {
	a := uint32(00000000000000000000000000001011)
	fmt.Println(hammingWeight(a))
}