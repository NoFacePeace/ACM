package array

import "sort"

func minTaps(n int, ranges []int) int {
	type tap struct {
		left  int
		right int
	}
	taps := []tap{}
	for i := 0; i <= n; i++ {
		t := tap{i - ranges[i], i + ranges[i]}
		taps = append(taps, t)
	}
	sort.Slice(taps, func(a, b int) bool {
		return taps[a].right < taps[b].right
	})
	start := 0
	i := -1
	cnt := 0
	for {
		old := i
		for j := i + 1; j < len(taps); j++ {
			if taps[j].left <= start {
				i = j
			}
		}
		if old == i {
			return -1
		}
		if taps[i].left > start {
			return -1
		}
		cnt++
		if taps[i].right >= n {
			return cnt
		}
		start = taps[i].right
	}

}
