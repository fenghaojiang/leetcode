func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}
	res := make([]int, len(queries))
	for k, subArray := range queries {
		res[k] = xors[subArray[0]] ^ xors[subArray[1]+1]
	}
	return res
}