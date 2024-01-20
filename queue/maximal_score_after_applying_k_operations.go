package queue

import "container/heap"

func maxKelements(nums []int, k int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	hp := &MaxHeapInt{}
	for _, v := range nums {
		heap.Push(hp, v)
	}
	i := 0
	ans := 0
	for i < k {
		i++
		val := heap.Pop(hp).(int)
		ans += val
		if val%3 != 0 {
			val += 3
		}
		heap.Push(hp, val/3)
	}
	return int64(ans)
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
