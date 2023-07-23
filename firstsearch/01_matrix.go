package firstsearch

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	if m == 0 {
		return mat
	}
	n := len(mat[0])
	if n == 0 {
		return mat
	}
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][j] = 1
				q = append(q, []int{i, j})
			}
		}
	}
	cnt := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] = 0
			arr[i][j] = cnt
			if i-1 >= 0 && mat[i-1][j] != 0 {
				q = append(q, []int{i - 1, j})
			}
			if i+1 < m && mat[i+1][j] != 0 {
				q = append(q, []int{i + 1, j})
			}
			if j-1 >= 0 && mat[i][j-1] != 0 {
				q = append(q, []int{i, j - 1})
			}
			if j+1 < n && mat[i][j+1] != 0 {
				q = append(q, []int{i, j + 1})
			}
		}
		cnt++
	}
	return arr
}
