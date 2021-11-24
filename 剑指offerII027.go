/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    mid := middleNode(head)
    l2 := mid.Next
    mid.Next = nil
    r2 := reverseList(l2)
    p2 := r2
    for head != nil && p2 != nil {
        if head.Val != p2.Val {
            return false
        }
        head = head.Next
        p2 = p2.Next
    }
    mid.Next = reverseList(r2)
    return true
}


func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
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


