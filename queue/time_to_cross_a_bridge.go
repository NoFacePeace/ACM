package queue

import (
	"container/heap"
	"math"
)

func findCrossingTime(n int, k int, time [][]int) int {
	rwk := &workQ{}
	lwk := &workQ{}
	rwt := &waitQ{}
	lwt := &waitQ{}
	for i := 0; i < len(time); i++ {
		heap.Push(lwt, wait{
			id:  i,
			eff: time[i][0] + time[i][2],
		})
	}
	cur := 0
	for n > 0 || rwk.Len() > 0 || rwt.Len() > 0 {
		for rwk.Len() > 0 && rwk.Top().(work).end <= cur {
			x := heap.Pop(rwk).(work)
			heap.Push(rwt, wait{
				id:  x.id,
				eff: time[x.id][0] + time[x.id][2],
			})
		}
		for lwk.Len() > 0 && lwk.Top().(work).end <= cur {
			x := heap.Pop(lwk).(work)
			heap.Push(lwt, wait{
				id:  x.id,
				eff: time[x.id][0] + time[x.id][2],
			})
		}
		if rwt.Len() > 0 {
			x := heap.Pop(rwt).(wait)
			cur += time[x.id][2]
			heap.Push(lwk, work{
				id:  x.id,
				end: cur + time[x.id][3],
			})
			continue
		}
		if n > 0 && lwt.Len() > 0 {
			x := heap.Pop(lwt).(wait)
			cur += time[x.id][0]
			heap.Push(rwk, work{
				id:  x.id,
				end: cur + time[x.id][1],
			})
			n--
			continue
		}
		nt := math.MaxInt
		if lwk.Len() > 0 {
			lt := lwk.Top().(work).end
			if lt < nt {
				nt = lt
			}
		}
		if rwk.Len() > 0 {
			rt := rwk.Top().(work).end
			if rt < nt {
				nt = rt
			}
		}
		if nt == math.MaxInt {
			continue
		}
		if nt > cur {
			cur = nt
		}
	}
	return cur
}

type work struct {
	id  int
	end int
}

type workQ []work

func (w workQ) Len() int {
	return len(w)
}

func (w workQ) Less(a, b int) bool {
	return w[a].end < w[b].end
}

func (w workQ) Swap(a, b int) {
	w[a], w[b] = w[b], w[a]
}

func (w workQ) Top() interface{} {
	return w[0]
}

func (w *workQ) Push(x interface{}) {
	*w = append(*w, x.(work))
}

func (w *workQ) Pop() interface{} {
	n := len(*w)
	x := (*w)[n-1]
	*w = (*w)[:n-1]
	return x
}

type wait struct {
	id  int
	eff int
}

type waitQ []wait

func (w waitQ) Len() int {
	return len(w)
}

func (w waitQ) Less(a, b int) bool {
	if w[a].eff == w[b].eff {
		return w[a].id > w[b].id
	}
	return w[a].eff > w[b].eff
}

func (w waitQ) Swap(a, b int) {
	w[a], w[b] = w[b], w[a]
}

func (w waitQ) Top() interface{} {
	return w[0]
}

func (w *waitQ) Push(x interface{}) {
	*w = append(*w, x.(wait))
}

func (w *waitQ) Pop() interface{} {
	n := len(*w)
	x := (*w)[n-1]
	*w = (*w)[:n-1]
	return x
}
