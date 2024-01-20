package queue

import "container/heap"

func halveArray(nums []int) int {
	q := &pqFloat{}
	sum := 0.0
	for _, v := range nums {
		f := float64(v)
		sum += f
		heap.Push(q, f)
	}
	half := 0.0
	cnt := 0
	for half*2 < sum {
		f := heap.Pop(q).(float64)
		half += f / 2
		cnt++
		heap.Push(q, f/2)
	}
	return cnt
}

type pqFloat []float64

func (q pqFloat) Len() int {
	return len(q)
}

func (q pqFloat) Less(a, b int) bool {
	return q[a] > q[b]
}

func (q pqFloat) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func (q *pqFloat) Push(x interface{}) {
	*q = append(*q, x.(float64))
}

func (q *pqFloat) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
