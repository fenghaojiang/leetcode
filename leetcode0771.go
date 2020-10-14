package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	m := make(map[byte]struct{})
	for i := range J {
		m[J[i]] = struct{}{}
	}
	cnt := 0
	for i := range S {
		if _, ok := m[S[i]]; ok {
			cnt++
		}
	}
	return cnt
}

func main() {
	J, S :=  "aA", "aAAbbbb"
	fmt.Println(numJewelsInStones(J,S))
}