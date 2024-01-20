package hash

func firstCompleteIndex(arr []int, mat [][]int) int {
	mp := map[int][]int{}
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[mat[i][j]] = []int{i, j}
		}
	}
	rows := make([]int, m)
	columns := make([]int, n)
	for i := 0; i < len(arr); i++ {
		x, y := mp[arr[i]][0], mp[arr[i]][1]
		rows[x]++
		columns[y]++
		if rows[x] == n || columns[y] == m {
			return i
		}
	}
	return -1
}
