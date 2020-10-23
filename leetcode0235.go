/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */


func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
    if node.Val < p.Val && node.Val < q.Val {
        return lowestCommonAncestor(node.Right, p, q)
    }
    if node.Val > p.Val && node.Val > q.Val {
        return lowestCommonAncestor(node.Left, p, q)
    }
    return node
}