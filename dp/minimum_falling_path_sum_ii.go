package dp

func minFallingPathSumII(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return grid[0][0]
	}
	one := 0
	two := 1
	for i := 1; i < n; i++ {
		if grid[0][i] < grid[0][two] {
			two = i
		}
		if grid[0][one] > grid[0][two] {
			one, two = two, one
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != one {
				grid[i][j] += grid[i-1][one]
			} else {
				grid[i][j] += grid[i-1][two]
			}
		}
		o := 0
		t := 1
		for j := 1; j < n; j++ {
			if grid[i][j] < grid[i][t] {
				t = j
			}
			if grid[i][o] > grid[i][t] {
				o, t = t, o
			}
		}
		one = o
		two = t
	}
	return grid[n-1][one]
}
