package queue

import "container/heap"

func minimumTime(n int, relations [][]int, time []int) int {
	ad := make([][]int, n+1)
	m := map[int]int{}
	for _, r := range relations {
		v1, v2 := r[0], r[1]
		ad[v1] = append(ad[v1], v2)
		m[v2]++
	}
	mt := map[int]int{}
	for i := 0; i < n; i++ {
		mt[i+1] = time[i]
	}
	q := &pqCourse{}
	for i := 1; i <= n; i++ {
		if m[i] == 0 {
			heap.Push(q, course{
				id:    i,
				month: mt[i],
			})
		}
	}
	month := 0
	for q.Len() > 0 {
		x := heap.Pop(q).(course)
		month = x.month
		for _, v := range ad[x.id] {
			m[v]--
			if m[v] == 0 {
				heap.Push(q, course{
					id:    v,
					month: x.month + mt[v],
				})
			}
		}
	}
	return month
}

type course struct {
	id    int
	month int
}
type pqCourse []course

func (q pqCourse) Len() int {
	return len(q)
}

func (q pqCourse) Less(a, b int) bool {
	return q[a].month < q[b].month
}

func (q pqCourse) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func (q *pqCourse) Push(x interface{}) {
	*q = append(*q, x.(course))
}

func (q *pqCourse) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
