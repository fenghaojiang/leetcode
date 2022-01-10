func validateStackSequences(pushed []int, popped []int) bool {
	var index int
	stack := []int{}
	for i := range pushed {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && stack[len(stack)-1] != popped[index] {
			index++
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}


