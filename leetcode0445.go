/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    r1 := reverseList(l1)
    r2 := reverseList(l2)
    dummy := new(ListNode)
    cur := dummy
    var carry int
    for r1 != nil || r2 != nil {
        tempNode := &ListNode{Val: carry}
        if r1 != nil {
            tempNode.Val += r1.Val
            r1 = r1.Next
        }
        if r2 != nil {
            tempNode.Val += r2.Val
            r2 = r2.Next
        }
        carry = tempNode.Val / 10
        tempNode.Val = tempNode.Val % 10
        cur.Next = tempNode
        cur = cur.Next
    }
    if carry > 0 {
        tempNode := &ListNode{Val: carry}
        cur.Next = tempNode
    }
    return reverseList(dummy.Next)
}

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