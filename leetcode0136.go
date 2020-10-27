package main

import "fmt"

func singleNumber(nums []int) int {
    single := 0
    for _, num := range nums {
        single ^= num
    }
    return single
}

func main() {
	a := []int{4,1,2,1,2}
	fmt.Println(singleNumber(a))
}