func findCenter(edges [][]int) int {
	points := make(map[int]int)
	for _, edg := range edges {
		points[edg[0]]++
		points[edg[1]]++
	}
	for key, p := range points {
		if p == len(edges) {
			return key
		}
	}
	return -1
}