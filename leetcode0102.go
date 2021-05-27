/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var dfs func(root *TreeNode, depth int)
	var mapDepth = make(map[int][]int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		mapDepth[depth] = append(mapDepth[depth], root.Val)
		if root.Left != nil {
			dfs(root.Left, depth+1)
		}
		if root.Right != nil {
			dfs(root.Right, depth+1)
		}
	}
	dfs(root, 0)
	res := make([][]int, 0)
	for i := 0; i < len(mapDepth); i++ {
		res = append(res, mapDepth[i])
	}
	return res
}

// 队列写法

func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

