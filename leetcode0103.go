/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	cur := []*TreeNode{root}
	res := make([][]int, 0)
	for level := 0; len(cur) > 0; level++ {
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
		if level%2 == 1 {
			for j := 0; j < len(temp)/2; j++ {
				temp[j], temp[len(temp)-j-1] = temp[len(temp)-j-1], temp[j]
			}
		}
		cur = cur[l:]
		res = append(res, temp)
	}
	return res
}