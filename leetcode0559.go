/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil || len(root.Children) == 0 {
		return 1
	}
	var max int = 0
	for i := 0; i < len(root.Children); i++ {
		temp := maxDepth(root.Children[i]) + 1
		if temp > max {
			max = temp
		}
	}
	return max
}