package main

import "fmt"


func isValid(s string) bool {
	if len(s)%2 == 1{
		return false
	}
	stack := make([]rune, 0)
	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
		}
		if r == '}' {
			if len(stack) != 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if r == ']' {
			if len(stack) != 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if r == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false 
	} else {
		return true
	}
}

func main() {
	fmt.Println(isValid("([)]"))
}
