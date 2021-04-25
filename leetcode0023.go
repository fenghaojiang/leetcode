
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	h := new(minHeap)
	for i := 0; i < l; i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		temp := heap.Pop(h).(*ListNode)
		if temp.Next != nil {
			heap.Push(h, temp.Next)
		}
		pre.Next = temp
		pre = pre.Next
	}

	return dummyHead.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
