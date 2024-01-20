package pointer

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	n := len(intervals)
	i := 0
	j := i + 1
	ans := [][]int{}
	for j < n {
		if intervals[i][1] < intervals[j][0] {
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
			i = j
			j = i + 1
			continue
		}
		if intervals[i][1] < intervals[j][1] {
			intervals[i][1] = intervals[j][1]
		}
		j++
	}
	ans = append(ans, []int{intervals[i][0], intervals[i][1]})
	return ans
}
