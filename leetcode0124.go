/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    var res int = math.MinInt32
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := max(dfs(root.Left), 0)
        right := max(dfs(root.Right), 0)
        temp := root.Val + left + right
        if res < temp {
            res = temp
        }
        return root.Val + max(left, right)
    }
    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}