package firstsearch

import "container/heap"

func minimumEffortPath(heights [][]int) int {
	row := len(heights)
	if row == 0 {
		return 0
	}
	column := len(heights[0])
	if column == 0 {
		return 0
	}
	if row == 1 && column == 1 {
		return 0
	}
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	q := &queue{}
	m := map[int]bool{}
	heap.Push(q, [3]int{0, 0, -1})
	for q.Len() > 0 {
		v := heap.Pop(q).([3]int)
		x, y, c := v[0], v[1], v[2]
		if x == row-1 && y == column-1 {
			return c
		}
		m[x*column+y] = true
		if x+1 < row && !m[(x+1)*column+y] {
			tmp := c
			if tmp == -1 {
				tmp = abs(heights[x][y], heights[x+1][y])
			} else {
				tmp = max(tmp, abs(heights[x][y], heights[x+1][y]))
			}
			heap.Push(q, [3]int{x + 1, y, tmp})
		}
		if x-1 >= 0 && !m[(x-1)*column+y] {
			tmp := c
			if tmp == -1 {
				tmp = abs(heights[x][y], heights[x-1][y])
			} else {
				tmp = max(tmp, abs(heights[x][y], heights[x-1][y]))
			}
			heap.Push(q, [3]int{x - 1, y, tmp})
		}
		if y+1 < column && !m[x*column+y+1] {
			tmp := c
			if tmp == -1 {
				tmp = abs(heights[x][y], heights[x][y+1])
			} else {
				tmp = max(tmp, abs(heights[x][y], heights[x][y+1]))
			}
			heap.Push(q, [3]int{x, y + 1, tmp})
		}
		if y-1 >= 0 && !m[x*column+y-1] {
			tmp := c
			if tmp == -1 {
				tmp = abs(heights[x][y], heights[x][y-1])
			} else {
				tmp = max(tmp, abs(heights[x][y], heights[x][y-1]))
			}
			heap.Push(q, [3]int{x, y - 1, tmp})
		}
	}
	return -1
}

type queue [][3]int

func (q queue) Len() int {
	return len(q)
}
func (q queue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func (q queue) Less(a, b int) bool {
	return q[a][2] < q[b][2]
}

func (q *queue) Push(x interface{}) {
	*q = append(*q, x.([3]int))
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
