/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(root *TreeNode)
	nodes := []int{}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			nodes = append(nodes, root.Val)
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root1)
	nodes1 := append([]int{}, nodes...)
	nodes = []int{}
	dfs(root2)
	if len(nodes) != len(nodes1) {
		return false
	}
	for i, v := range nodes {
		if v != nodes1[i] {
			return false
		}
	}
	return true
}