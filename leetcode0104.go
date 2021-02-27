/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		return maxDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return maxDepth(root.Right) + 1
	} else {
		return max(maxDepth(root.Right)+1, maxDepth(root.Left)+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}