/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cur := head
    for cur.Next != nil {
        if cur.Val != cur.Next.Val {
            cur = cur.Next
        } else {
            cur.Next = cur.Next.Next
        }
    }
    return head
}

