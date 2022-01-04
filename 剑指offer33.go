func verifyPostorder(postorder []int) bool {
	var helper func(postorder []int, left, right int) bool
	helper = func(postorder []int, left, right int) bool {
		if left >= right {
			return true
		}
		p := left
		for postorder[p] < postorder[right] {
			p++
		}
		m := p
		for postorder[p] > postorder[right] {
			p++
		}
		return p == right && helper(postorder, left, m-1) && helper(postorder, m, right-1)
	}
	return helper(postorder, 0, len(postorder)-1)
}