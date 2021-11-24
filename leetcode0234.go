package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

func main() {
	a := []int{1, 2, 2, 1}
	head := new(ListNode)
	pre := head
	for i := range a {
		node := new(ListNode)
		node.Val = a[i]
		pre.Next = node
		pre = node
	}
	if pre == nil {

	}
	fmt.Println(isPalindrome(head.Next))
}



/////v2



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
    r2 := reverse(l2)
    p2 := r2
    for head != nil && p2 != nil {
        if head.Val != p2.Val {
            return false
        }
        head = head.Next
        p2 = p2.Next
    }
    mid.Next = reverse(r2)
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


func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}