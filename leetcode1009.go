package main

import "fmt"

func bitwiseComplement(N int) int {
    if N == 0 {
        return 1
    }
    
    temp := 1
    for i := N; i > 0; i /= 2 {
        temp <<= 1
    }
    
    return (temp-1)^N
}

func main() {
	
	fmt.Println(bitwiseComplement(5))
}