/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findTilt(root *TreeNode) int {
    var res int
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        leftSum := dfs(root.Left)
        rightSum := dfs(root.Right)
        res += abs(leftSum - rightSum)
        return leftSum + rightSum + root.Val
    }
    dfs(root)
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}


