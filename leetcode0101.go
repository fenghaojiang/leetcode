/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    return check(root.Left, root.Right)
}


func check(p, q *TreeNode) bool {
    if q == nil && p == nil {
        return true
    }
    if q == nil || p == nil {
        return false
    }
    return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}