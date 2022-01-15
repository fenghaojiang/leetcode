const modNum = 1000000007

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	k, v := n/3, n%3
	if v == 0 {
		return Pow(3, k)
	} else if v == 1 {
		return Pow(3, k-1) * 4 % modNum
	}
	return Pow(3, k) * 2 % modNum
}
func Pow(x int, y int) int { //快速幂算法
	var result = 1
	for y > 0 {
		if y%2 == 1 {
			result = result * x % modNum
		}
		y = y / 2
		x = x * x % modNum
	}
	return result
}
