func decode(encoded []int, first int) []int {
	res := make([]int, 0)
	res = append(res, first)
	for i := range encoded {
		res = append(res, encoded[i]^res[i])
	}
	return res
}