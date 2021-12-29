func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var isSameBias func(arr []int) bool
	isSameBias = func(arr []int) bool {
		sort.Ints(arr)
		var bias int
		for i := 1; i < len(arr); i++ {
			if i == 1 {
				bias = arr[i] - arr[i-1]
			} else if bias != arr[i]-arr[i-1] {
				return false
			}
		}
		return true
	}
	res := make([]bool, 0)
	for i := range l {
		res = append(res, isSameBias(append([]int{}, nums[l[i]:r[i]+1]...)))
	}
	return res
}

