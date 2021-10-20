/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	rightNode := prev
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := prev.Next
	tail := rightNode.Next
	prev.Next = nil
	rightNode.Next = nil
	reverseInside(leftNode)
	prev.Next = rightNode
	leftNode.Next = tail
	return dummy.Next
}

func reverseInside(head *ListNode) {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}