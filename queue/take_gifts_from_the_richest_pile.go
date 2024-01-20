package queue

import "container/heap"

func pickGifts(gifts []int, k int) int64 {
	tmp := MaxHeapInt(gifts)
	h := &tmp
	heap.Init(h)
	i := 0
	sqr := func(num int) int {
		i := 0
		for i*i <= num {
			i++
		}
		return i - 1
	}
	for i < k {
		i++
		x := heap.Pop(h).(int)
		heap.Push(h, sqr(x))
	}
	ans := 0
	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		ans += x
	}
	return int64(ans)
}
