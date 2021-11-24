/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeList(l1, l2 *ListNode) {
	var l1next, l2next *ListNode
	for l1 != nil && l2 != nil {
		l1next = l1.Next
		l2next = l2.Next

		l1.Next = l2
		l1 = l1next

		l2.Next = l1
		l2 = l2next
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l2 := mid.Next
	mid.Next = nil
	r2 := reverseList(l2)
	mergeList(head, r2)
}