var (
	src string
	ptr int
)

func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}

func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	times := 1
	res := ""
	if cur >= '0' && cur <= '9' {
		times = getDigit()
		ptr++
		str := getString()
		ptr++
		res = strings.Repeat(str, times)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		res = string(cur)
		ptr++
	}
	return res + getString()
}

func getDigit() int {
	res := 0
	for src[ptr] >= '0' && src[ptr] <= '9' {
		res = res*10 + int(src[ptr]-'0')
		ptr++
	}
	return res
}