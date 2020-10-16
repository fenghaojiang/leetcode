package main

import "fmt"

func fib(N int) int {
    if N == 0 {
        return 0
    } else if N == 1 {
        return 1
    }
    a := 0
    b := 1
    var c int
    for i := 1; i < N; i++ {
        c = a + b
        a = b 
        b = c
    }
    return c
}

func main() {
	N := 4
	fmt.Println(fib(N))
}