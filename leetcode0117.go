/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := []*Node{root}
	for len(cur) != 0 {
		l := len(cur)
		for i := 0; i < l; i++ {
			if i == l-1 {
				cur[i].Next = nil
			} else {
				cur[i].Next = cur[i+1]
			}
			if cur[i].Left != nil {
				cur = append(cur, cur[i].Left)
			}
			if cur[i].Right != nil {
				cur = append(cur, cur[i].Right)
			}
		}
		cur = cur[l:]
	}
	return root
}