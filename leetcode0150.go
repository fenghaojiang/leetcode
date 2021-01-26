package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	for _, v := range tokens {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		} else {
			num1 := nums[len(nums)-1]
			num2 := nums[len(nums)-2]
			nums[len(nums)-2] = cal(num2, num1, v)
			nums = nums[:len(nums)-1]
		}
	}
	return nums[0]
}

func cal(a, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result := evalRPN(tokens)
	fmt.Println(result)
}
