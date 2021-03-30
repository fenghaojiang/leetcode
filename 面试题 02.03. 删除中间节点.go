/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(node *ListNode) {
    for node != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
		}
		node = node.Next
	}
}