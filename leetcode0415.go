func addStrings(num1 string, num2 string) string {
	carry := 0
	res := ""
	for i, j := len(num1)-1, len(num2)-1; j >= 0 || i >= 0 || carry != 0; i, j = i-1, j-1 {
		cal := carry
		if i >= 0 {
			cal += int(num1[i] - '0')
		}
		if j >= 0 {
			cal += int(num2[j] - '0')
		}
		if cal >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		res = strconv.Itoa(cal%10) + res
	}
	return res
} 