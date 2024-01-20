package greedy

import "container/heap"

func minStoneSum(piles []int, k int) int {
	q := &MaxHeapInt{}
	sum := 0
	for _, v := range piles {
		sum += v
		heap.Push(q, v)
	}
	for i := 0; i < k; i++ {
		if q.Len() == 0 {
			break
		}
		x := heap.Pop(q).(int)
		sum -= x
		sum += (x + 1) / 2
		heap.Push(q, (x+1)/2)
	}
	return sum
}
