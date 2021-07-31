func getLucky(s string, k int) int {
	sum := 0
	for _, b := range s {
		b -= 'a' - 1
		sum += int(b/10 + b%10)
	}
	for k--; k > 0 && sum > 9; k-- {
		s := sum
		for sum = 0; s > 0; s /= 10 {
			sum += s % 10
		}
	}
	return sum
}