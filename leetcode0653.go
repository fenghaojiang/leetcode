/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	var dfs func(root *TreeNode, k int) bool
	dfs = func(root *TreeNode, k int) bool {
		if root == nil {
			return false
		}
		if _, ok := m[k-root.Val]; ok {
			return true
		} else {
			m[root.Val] = struct{}{}
		}
		return dfs(root.Left, k) || dfs(root.Right, k)
	}
	return dfs(root, k)
}