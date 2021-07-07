/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := dfs(start, i-1)
		right := dfs(i+1, end)
		for _, leftTree := range left {
			for _, rightTree := range right {
				root := &TreeNode{i, nil, nil}
				root.Left = leftTree
				root.Right = rightTree
				res = append(res, root)
			}
		}
	}
	return res
}

