package main

import "fmt"


var res int

func sumNums(n int) int {
	res = 0
	sum(n)
	return res
}

func sum(n int) bool {
	res += n
	return 0 < n && sum(n-1)
}

func main() {
	n := 9
	fmt.Println(sumNums(n))
}