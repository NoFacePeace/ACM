package firstsearch

import "sort"

func pondSizes(land [][]int) []int {
	n := len(land)
	if n == 0 {
		return nil
	}
	m := len(land[0])
	if m == 0 {
		return nil
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 {
			return 0
		}
		if y < 0 {
			return 0
		}
		if x >= n {
			return 0
		}
		if y >= m {
			return 0
		}
		if land[x][y] != 0 {
			return 0
		}
		land[x][y] = 1
		cnt := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				cnt += dfs(x+i, y+j)
			}
		}
		return cnt + 1
	}
	arr := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] != 0 {
				continue
			}
			cnt := dfs(i, j)
			arr = append(arr, cnt)
		}
	}
	sort.Ints(arr)
	return arr
}
