/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rootSum(root *TreeNode, targetSum int) int {
	var res int
	if root == nil {
		return res
	}
	if targetSum == root.Val {
		res++
	}
	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	preSum := make(map[int]int)
	preSum[0] = 1
	var dfs func(node *TreeNode, cur int)
	var res int
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		res += preSum[cur-targetSum]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
	}
	dfs(root, 0)
	return res
}
