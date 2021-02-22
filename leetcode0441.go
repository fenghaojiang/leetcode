package main

func main() {
	fmt.Println(arrangeCoins(5))
}

func arrangeCoins(n int) int {
	var line int = 1
	count := 0
	for n >= 0 {
		if n-line >= 0 {
			count++
		}
		n = n - line
		line++
	}
	return count
}
