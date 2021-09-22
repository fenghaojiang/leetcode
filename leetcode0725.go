/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	var l int
	for cur := head; cur != nil; cur = cur.Next {
		l++
	}
	size, remain := l/k, l%k
	res := make([]*ListNode, k)
	for i, cur := 0, head; i < k && cur != nil; i++ {
		res[i] = cur
		curSize := size
		if i < remain {
			curSize++
		}
		for j := 1; j < curSize; j++ {
			cur = cur.Next
		}
		cur, cur.Next = cur.Next, nil
	}
	return res
}