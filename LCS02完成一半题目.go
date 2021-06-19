func halfQuestions(questions []int) int {
	arr := make([]int, 1001)
	ans := 0
	for _, v := range questions {
		arr[v]++
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	n := len(questions) / 2
	for n > 0 {
		n -= arr[ans]
		ans++
	}
	return ans
}