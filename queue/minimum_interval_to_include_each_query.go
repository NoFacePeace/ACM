package queue

import (
	"container/heap"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	sq := [][]int{}
	for i := 0; i < len(queries); i++ {
		sq = append(sq, []int{i, queries[i]})
	}
	sort.Slice(sq, func(a, b int) bool {
		return sq[a][1] < sq[b][1]
	})
	ans := make([]int, len(queries))
	i := 0
	q := &pqq{}
	for j := 0; j < len(sq); j++ {
		idx, val := sq[j][0], sq[j][1]
		for i < len(intervals) {
			if intervals[i][0] > val {
				break
			}
			heap.Push(q, intervals[i])
			i++
		}
		for len(*q) > 0 {
			x := heap.Pop(q).([]int)
			if x[1] >= val {
				ans[idx] = x[1] - x[0] + 1
				heap.Push(q, x)
				break
			}
		}
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] == 0 {
			ans[i] = -1
		}
	}
	return ans
}

type pqq [][]int

func (q pqq) Len() int {
	return len(q)
}

func (q pqq) Less(a, b int) bool {
	return q[a][1]-q[a][0] < q[b][1]-q[b][0]
}

func (q pqq) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func (q *pqq) Push(x interface{}) {
	*q = append(*q, x.([]int))
}

func (q *pqq) Pop() interface{} {
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}
