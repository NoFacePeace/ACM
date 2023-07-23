package math

func nearestValidPoint(x int, y int, points [][]int) int {
	min := -1
	ans := -1
	xy := -1
	for i := 0; i < len(points); i++ {
		p := points[i]
		if p[0] != x && p[1] != y {
			continue
		}
		val := abs(x, p[0]) + abs(y, p[1])
		if min == -1 {
			min = val
			ans = i
			xy = x + y
			continue
		}
		if val > min {
			continue
		}
		if val == min {
			if p[0]+p[1] >= xy {
				continue
			}
			xy = p[0] + p[1]
			ans = i
			continue
		}
		min = val
		xy = p[0] + p[1]
		ans = i
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
