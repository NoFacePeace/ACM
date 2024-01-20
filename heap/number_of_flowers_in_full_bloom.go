package heap

import (
	"container/heap"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	startHeap := &MinHeapInt{}
	endHeap := &MinHeapInt{}
	for _, flower := range flowers {
		start, end := flower[0], flower[1]
		heap.Push(startHeap, start)
		heap.Push(endHeap, end)
	}
	type p struct {
		id int
		t  int
	}
	arr := []p{}
	for k, v := range people {
		arr = append(arr, p{id: k, t: v})
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a].t < arr[b].t
	})
	cnt := 0
	n := len(people)
	ans := make([]int, n)
	for _, a := range arr {
		id, v := a.id, a.t
		for startHeap.Len() > 0 && startHeap.Top() <= v {
			cnt++
			heap.Pop(startHeap)
		}
		for endHeap.Len() > 0 && endHeap.Top() < v {
			cnt--
			heap.Pop(endHeap)
		}
		ans[id] = cnt
	}
	return ans
}

type MinHeapInt []int

func (h MinHeapInt) Len() int {
	return len(h)
}

func (h MinHeapInt) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h MinHeapInt) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeapInt) Top() int {
	return h[0]
}
