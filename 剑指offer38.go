func permutation(s string) []string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	l := len(arr)
	visit := make([]bool, l)
	res := make([]string, 0)
	temp := make([]byte, 0)
	var backTrack func(index int)
	backTrack = func(index int) {
		if index == l {
			res = append(res, string(temp))
			return
		}
		for i, v := range visit {
			if v || (i > 0 && !visit[i-1] && arr[i-1] == arr[i]) {
				continue
			}
			visit[i] = true
			temp = append(temp, arr[i])
			backTrack(index + 1)
			temp = temp[:len(temp)-1]
			visit[i] = false
		}
	}
	backTrack(0)
	return res
}