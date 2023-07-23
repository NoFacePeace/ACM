package queue

import (
	"container/heap"
	"math"
)

func findMaxValueOfEquation(points [][]int, k int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	q := &pq{}
	heap.Push(q, point{
		x: points[0][0],
		y: points[0][1],
		w: points[0][1] - points[0][0],
	})
	ans := math.MinInt
	for i := 1; i < n; i++ {
		p := points[i]

		for q.Len() > 0 && p[0]-q.Top().x > k {
			heap.Pop(q)
		}
		if q.Len() > 0 {
			t := q.Top()
			dst := p[0] - t.x + p[1] + t.y
			if dst > ans {
				ans = dst
			}
		}
		heap.Push(q, point{
			x: p[0],
			y: p[1],
			w: p[1] - p[0],
		})
	}
	return ans
}

type point struct {
	x int
	y int
	w int
}
type pq []point

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(a, b int) bool {
	return q[a].w > q[b].w
}

func (q pq) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func (q *pq) Push(x interface{}) {
	*q = append(*q, x.(point))
}

func (q *pq) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q pq) Top() point {
	return q[0]
}
