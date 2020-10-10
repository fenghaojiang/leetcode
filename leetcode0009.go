package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0){
		return false
	} 
	tenetNumber := 0
	for x > tenetNumber {
		tenetNumber = tenetNumber*10 + x%10
		x /= 10
	}
	return x == tenetNumber || x == tenetNumber / 10 //奇数的情况
}

func main() {
	fmt.Println(isPalindrome(121))
}