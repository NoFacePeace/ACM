package firstsearch

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	l := len(jobDifficulty)
	if l <= 0 {
		return -1
	}
	if l < d {
		return -1
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, d+1)
	}
	arr := make([][]int, l)
	for i := 0; i < l; i++ {
		arr[i] = make([]int, l)
		max := jobDifficulty[i]
		for j := i; j < l; j++ {
			if jobDifficulty[j] > max {
				max = jobDifficulty[j]
			}
			arr[i][j] = max
		}
	}
	if arr[0][l-1] == 0 {
		return 0
	}
	var f func(idx, day int) int
	f = func(idx, day int) int {
		if day == 1 {
			return arr[idx][l-1]
		}
		if dp[idx][day] != 0 {
			return dp[idx][day]
		}
		min := math.MaxInt
		for i := 0; i <= l-d; i++ {
			if l-idx-i < day {
				break
			}
			val := arr[idx][idx+i] + f(idx+i+1, day-1)
			if val < min {
				min = val
			}
		}
		dp[idx][day] = min
		return min
	}
	return f(0, d)
}
