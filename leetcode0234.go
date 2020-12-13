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
