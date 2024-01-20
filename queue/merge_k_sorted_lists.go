package queue

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	q := &listQ{}
	for _, v := range lists {
		if v == nil {
			continue
		}
		heap.Push(q, v)
	}
	h := &ListNode{}
	next := h
	for q.Len() > 0 {
		x := heap.Pop(q).(*ListNode)
		next.Next = x
		next = next.Next
		x = x.Next
		if x != nil {
			heap.Push(q, x)
		}
	}
	return h.Next
}

type listQ []*ListNode

func (l listQ) Len() int {
	return len(l)
}

func (l listQ) Swap(a, b int) {
	l[a], l[b] = l[b], l[a]
}

func (l listQ) Less(a, b int) bool {
	return l[a].Val < l[b].Val
}

func (l *listQ) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

func (l *listQ) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[:n-1]
	return x
}
