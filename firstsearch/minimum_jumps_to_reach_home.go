package firstsearch

import "math"

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}
	m := map[int]bool{}
	for _, v := range forbidden {
		m[v] = true
	}
	max := 8001
	min := 0
	dp := make([][2]int, 8002)
	var f func(x, d, c int)
	f = func(x, d, c int) {
		if x < min {
			return
		}
		if x > max {
			return
		}
		if m[x] {
			return
		}
		if dp[x][d] != 0 {
			return
		}
		dp[x][d] = c
		f(x+a, 1, c+1)
		if d != 0 {
			f(x-b, 0, c+1)
		}
	}
	f(a, 1, 1)
	left := math.MaxInt
	if dp[x][0] != 0 {
		left = dp[x][0]
	}
	right := math.MaxInt
	if dp[x][1] != 0 {
		right = dp[x][1]
	}
	ans := left
	if ans > right {
		ans = right
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
