package array

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m := len(image)
	if m == 0 {
		return image
	}
	n := len(image[0])
	if n == 0 {
		return image
	}
	old := image[sr][sc]
	if old == color {
		return image
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 {
			return
		}
		if x >= m {
			return
		}
		if y < 0 {
			return
		}
		if y >= n {
			return
		}
		if image[x][y] != old {
			return
		}
		image[x][y] = color
		dfs(x, y-1)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x+1, y)
	}
	dfs(sr, sc)
	return image
}
