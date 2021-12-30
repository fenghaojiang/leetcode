/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	temp := []int{}
	var dfs func(root *TreeNode, remain int)
	dfs = func(root *TreeNode, remain int) {
		if root == nil {
			return
		}
		remain -= root.Val
		temp = append(temp, root.Val)
		if root.Left == nil && root.Right == nil && remain == 0 {
			res = append(res, append([]int{}, temp...))
		}
		defer func() { temp = temp[:len(temp)-1] }()
		dfs(root.Left, remain)
		dfs(root.Right, remain)
	}
	dfs(root, target)
	return res
}