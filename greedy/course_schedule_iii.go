package greedy

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(a, b int) bool {
		return courses[a][1] < courses[b][1]
	})
	h := &MaxHeapInt{}
	total := 0
	for _, course := range courses {
		t := course[0]
		ld := course[1]
		if t+total <= ld {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.Top() {
			x := heap.Pop(h).(int)
			total += t - x
			heap.Push(h, t)
		}
	}
	return h.Len()
}

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

func (h MaxHeapInt) Top() int {
	return h[0]
}
