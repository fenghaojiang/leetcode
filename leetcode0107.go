/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	cur := []*TreeNode{root}
	res := make([][]int, 0)
	for len(cur) > 0 {
		l := len(cur)
		temp := make([]int, 0)
		for i := range cur {
			temp = append(temp, cur[i].Val)
			if cur[i].Left != nil {
				cur = append(cur, cur[i].Left)
			}
			if cur[i].Right != nil {
				cur = append(cur, cur[i].Right)
			}
		}
		cur = cur[l:]
		res = append(res, temp)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}