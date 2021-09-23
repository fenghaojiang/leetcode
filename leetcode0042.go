func trap(height []int) int {
	start, end := 0, len(height)-1
	leftMax, rightMax := 0, 0
	var res int
	for start < end {
		leftMax = max(leftMax, height[start])
		rightMax = max(rightMax, height[end])
		if height[start] < height[end] {
			res += leftMax - height[start]
			start++
		} else {
			res += rightMax - height[end]
			end--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}