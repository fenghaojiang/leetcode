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
    res += rootSum(root.Left, targetSum - root.Val)
    res += rootSum(root.Right, targetSum - root.Val)
    return res
}


func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    return rootSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}