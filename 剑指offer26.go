/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubStructure(A *TreeNode, B *TreeNode) bool {
    return A != nil && B != nil && (isRecur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func isRecur(A, B *TreeNode) bool {
    if B == nil {
        return true
    }
    if A == nil || A.Val != B.Val {
        return false
    }
    return isRecur(A.Left, B.Left) && isRecur(A.Right, B.Right)
}