/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	max := -1
	maxIndex := -1
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	left := constructMaximumBinaryTree(nums[:maxIndex])
	right := constructMaximumBinaryTree(nums[maxIndex+1:])
	return &TreeNode{Val: max, Left: left, Right: right}
}