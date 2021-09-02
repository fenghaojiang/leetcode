/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var n int
	cur := new(ListNode)
	for i := head; i != nil; i = i.Next {
		n++
	}
	for cur = head; n > k; cur = cur.Next {
		n--
	}
	return cur
}