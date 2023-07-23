package array

import (
	"container/heap"
)

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	q := &ClassQueue{}
	heap.Init(q)
	for _, v := range classes {
		rate1 := float64(v[0]) / float64(v[1])
		rate2 := float64(v[0]+1) / float64(v[1]+1)
		dist := rate2 - rate1
		c := Class{v[0], v[1], dist}
		heap.Push(q, c)
	}
	i := 0
	for i < extraStudents {
		x := heap.Pop(q).(Class)
		x.pass++
		x.total++
		r1 := float64(x.pass) / float64(x.total)
		r2 := float64(x.pass+1) / float64(x.total+1)
		x.distance = r2 - r1
		heap.Push(q, x)
		i++
	}
	ans := 0.0
	for q.Len() > 0 {
		x := heap.Pop(q).(Class)
		ans += float64(x.pass) / float64(x.total)
	}
	return ans / float64(len(classes))
}

type Class struct {
	pass     int
	total    int
	distance float64
}

type ClassQueue []Class

func (c ClassQueue) Len() int {
	return len(c)
}

func (c ClassQueue) Less(i, j int) bool {
	return c[i].distance > c[j].distance
}

func (c ClassQueue) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *ClassQueue) Push(x interface{}) {
	*c = append(*c, x.(Class))
}

func (c *ClassQueue) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}
