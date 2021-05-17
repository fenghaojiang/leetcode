/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
		if xFound && yFound {
			return
		}
		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}