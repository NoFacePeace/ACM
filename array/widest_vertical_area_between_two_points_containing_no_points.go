package array

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(a, b int) bool {
		return points[a][0] < points[b][0]
	})
	max := 0
	for i := 0; i < len(points)-1; i++ {
		dist := points[i+1][0] - points[i][0]
		if dist > max {
			max = dist
		}
	}
	return max
}
