/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    curNode := root
    for curNode != nil {
        if curNode.Left != nil {
            next := curNode.Left
            p := next 
            for p.Right != nil {
                p = p.Right
            }
        
            p.Right = curNode.Right
            curNode.Left = nil
            curNode.Right = next
        }
        curNode = curNode.Right
    }
}