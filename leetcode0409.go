package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	res := 0 
	for _, c := range s {
		m[c]++
	}
	for _, value := range m {
		res += value/2 * 2
	}
	if len(s) > res {
		res++
	}
	return res
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}