/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	res := -1
	rootVal := root.Val
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || (res != -1 && node.Val >= res) {
			return
		}
		if node.Val > rootVal {
			res = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}