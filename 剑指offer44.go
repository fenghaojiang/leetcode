func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	num := start + (n-1)/digit
	return int(fmt.Sprintf("%d", num)[(n-1)%digit] - '0')
}

