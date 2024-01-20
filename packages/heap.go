package packages

type MaxHeapInt []int

func (h MaxHeapInt) Len() int {
	return len(h)
}

func (h MaxHeapInt) Less(a, b int) bool {
	return h[a] > h[b]
}

func (h MaxHeapInt) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *MaxHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
