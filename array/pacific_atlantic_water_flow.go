package array

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	if m == 0 {
		return nil
	}
	n := len(heights[0])
	if n == 0 {
		return nil
	}
	q := [][]int{}
	for i := 0; i < m; i++ {
		q = append(q, []int{i, 0})
	}
	for i := 1; i < n; i++ {
		q = append(q, []int{0, i})
	}
	visit := make([][]bool, m)
	p := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
		p[i] = make([]bool, n)
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			if visit[i][j] {
				continue
			}
			visit[i][j] = true
			p[i][j] = true
			if i-1 >= 0 && !visit[i-1][j] && heights[i-1][j] >= heights[i][j] {
				q = append(q, []int{i - 1, j})
			}
			if i+1 < m && !visit[i+1][j] && heights[i+1][j] >= heights[i][j] {
				q = append(q, []int{i + 1, j})
			}
			if j-1 >= 0 && !visit[i][j-1] && heights[i][j-1] >= heights[i][j] {
				q = append(q, []int{i, j - 1})
			}
			if j+1 < n && !visit[i][j+1] && heights[i][j+1] >= heights[i][j] {
				q = append(q, []int{i, j + 1})
			}
		}
	}
	a := make([][]bool, m)
	for i := 0; i < m; i++ {
		a[i] = make([]bool, n)
		visit[i] = make([]bool, n)
	}
	q = [][]int{}
	for i := 0; i < m; i++ {
		q = append(q, []int{i, n - 1})
	}
	for i := 0; i < n-1; i++ {
		q = append(q, []int{m - 1, i})
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			if visit[i][j] {
				continue
			}
			visit[i][j] = true
			a[i][j] = true
			if i-1 >= 0 && !visit[i-1][j] && heights[i-1][j] >= heights[i][j] {
				q = append(q, []int{i - 1, j})
			}
			if i+1 < m && !visit[i+1][j] && heights[i+1][j] >= heights[i][j] {
				q = append(q, []int{i + 1, j})
			}
			if j-1 >= 0 && !visit[i][j-1] && heights[i][j-1] >= heights[i][j] {
				q = append(q, []int{i, j - 1})
			}
			if j+1 < n && !visit[i][j+1] && heights[i][j+1] >= heights[i][j] {
				q = append(q, []int{i, j + 1})
			}
		}
	}
	ans := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
