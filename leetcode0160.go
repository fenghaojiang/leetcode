package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	for ha != hb {
		if ha != nil {
			ha = ha.Next
		} else {
			ha = headB
		}
		if hb != nil {
			hb = hb.Next
		} else {
			hb = headA
		}
	}
	return ha
}
