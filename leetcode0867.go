package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	b := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(a, b))
}

func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
