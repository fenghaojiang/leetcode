/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	temp := dummyHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return dummyHead.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}