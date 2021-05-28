/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return (abs(dfs(root.Left)-dfs(root.Right)) <= 1) && isBalanced(root.Right) && isBalanced(root.Left)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}