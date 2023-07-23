package algorithm

import "container/heap"

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	adList := map[int][][]int{}
	for i := 0; i < len(edges); i++ {
		u, v, w := edges[i][0], edges[i][1], edges[i][2]
		adList[u] = append(adList[u], []int{v, w})
		adList[v] = append(adList[v], []int{u, w})
	}
	h := &myHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, 0})
	visit := map[int]bool{}
	ans := 0
	used := make([]int, n*n)
	for h.Len() > 0 {
		arr := heap.Pop(h).([]int)
		u, step := arr[0], arr[1]
		if visit[u] {
			continue
		}
		visit[u] = true
		ans++
		for i := 0; i < len(adList[u]); i++ {
			v, nodes := adList[u][i][0], adList[u][i][1]
			if step+nodes+1 <= maxMoves && !visit[v] {
				heap.Push(h, []int{v, step + nodes + 1})
			}
			used[u*n+v] = nodes
			if used[u*n+v] > maxMoves-step {
				used[u*n+v] = maxMoves - step
			}
		}
	}
	for _, edge := range edges {
		u, v, nodes := edge[0], edge[1], edge[2]
		cnt := nodes
		if cnt > used[u*n+v]+used[v*n+u] {
			cnt = used[u*n+v] + used[v*n+u]
		}
		ans += cnt
	}
	return ans
}

type myHeap [][]int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Less(a, b int) bool {
	return h[a][1] < h[b][1]
}

func (h myHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *myHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
