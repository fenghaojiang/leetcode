package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	if l1 < l2 {
		return gcdOfStrings(str2, str1)
	}

	if l1 == l2 {
		if str1 == str2 {
			return str1
		}

		return ""
	}

	for i := 0; i < l2; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}

	return gcdOfStrings(str1[l2:], str2)
}

func main() {
	str1, str2 := "ABCABC", "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}