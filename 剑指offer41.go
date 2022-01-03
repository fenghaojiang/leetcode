type maxHeap []int // 大顶堆
type minHeap []int // 小顶堆

// 每个堆都要heap.Interface的五个方法：Len, Less, Swap, Push, Pop
// 其实只有Less的区别。

// Len 返回堆的大小
func (m maxHeap) Len() int {
	return len(m)
}
func (m minHeap) Len() int {
	return len(m)
}

// Less 决定是大优先还是小优先
func (m maxHeap) Less(i, j int) bool { // 大根堆
	return m[i] > m[j]
}
func (m minHeap) Less(i, j int) bool { // 小根堆
	return m[i] < m[j]
}

// Swap 交换下标i, j元素的顺序
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Push 在堆的末尾添加一个元素，注意和heap.Push(heap.Interface, interface{})区分
func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

// Pop 删除堆尾的元素，注意和heap.Pop()区分
func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}

// MedianFinder 维护两个堆，前一半是大顶堆，后一半是小顶堆，中位数由两个堆顶决定
type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

// Constructor 建两个空堆
func Constructor() MedianFinder {
	return MedianFinder{
		new(maxHeap),
		new(minHeap),
	}
}

// AddNum 插入元素num
// 分两种情况插入：
// 1. 两个堆的大小相等，则小顶堆增加一个元素（增加的不一定是num）
// 2. 小顶堆比大顶堆多一个元素，大顶堆增加一个元素
// 这两种情况又分别对应两种情况：
// 1. num小于大顶堆的堆顶，则num插入大顶堆
// 2. num大于小顶堆的堆顶，则num插入小顶堆
// 插入完成后记得调整堆的大小使得两个堆的容量相等，或小顶堆大1
func (m *MedianFinder) AddNum(num int) {
	if m.maxH.Len() == m.minH.Len() {
		if m.minH.Len() == 0 || num >= (*m.minH)[0] {
			heap.Push(m.minH, num)
		} else {
			heap.Push(m.maxH, num)
			top := heap.Pop(m.maxH).(int)
			heap.Push(m.minH, top)
		}
	} else {
		if num > (*m.minH)[0] {
			heap.Push(m.minH, num)
			bottle := heap.Pop(m.minH).(int)
			heap.Push(m.maxH, bottle)
		} else {
			heap.Push(m.maxH, num)
		}
	}
}

// FindMediam 输出中位数
func (m *MedianFinder) FindMedian() float64 {
	if m.minH.Len() == m.maxH.Len() {
		return float64((*m.maxH)[0])/2.0 + float64((*m.minH)[0])/2.0
	} else {
		return float64((*m.minH)[0])
	}
}