/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	curNode := []*TreeNode{root}
	for len(curNode) != 0 {
		l := len(curNode)
		for i := range curNode {
			if curNode[i] != nil {
				res = append(res, curNode[i].Val)
				curNode = append(curNode, curNode[i].Left)
				curNode = append(curNode, curNode[i].Right)
			}
		}
		//curNode = append([]*TreeNode{}, curNode[1:]...)
		curNode = curNode[l:]
	}
	return res
}