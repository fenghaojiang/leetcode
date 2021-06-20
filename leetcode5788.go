func largestOddNumber(num string) string {
	l := len(num)
	if l == 0 {
		return ""
	}
	for l != 0 && (num[l-1]-'0')%2 == 0 {
		num = num[:l-1]
		l = len(num)
	}
	return num
}