package hash

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	arr := []int{}
	for i := 0; i < m; i++ {
		num := 0
		for j := 0; j < n; j++ {
			num = num<<1 + matrix[i][j]
		}
		arr = append(arr, num)
	}
	max := 0
	for i := 0; i < m; i++ {
		cnt := 1
		for j := 0; j < m; j++ {
			tmp := arr[i] ^ arr[j]
			if tmp == (1<<n)-1 {
				cnt++
			}
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
