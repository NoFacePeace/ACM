package array

func kthSmallest(mat [][]int, k int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	arr := []int{}
	rerank := func(arr []int) {
		n := len(arr)
		for i := n - 1; i > 0; i-- {
			if arr[i] >= arr[i-1] {
				break
			}
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
	}
	for i := 0; i < m; i++ {
		n := len(mat[i])
		if i == 0 {
			for j := 0; j < n; j++ {
				if len(arr) >= k {
					break
				}
				arr = append(arr, mat[i][j])
			}
			continue
		}
		tmp := arr
		arr = nil
		for j := 0; j < len(tmp); j++ {
			for x := 0; x < n; x++ {
				if len(arr) < k {
					arr = append(arr, tmp[j]+mat[i][x])
					rerank(arr)
					continue
				}
				if arr[k-1] <= tmp[j]+mat[i][x] {
					n = x
					break
				}
				arr[k-1] = tmp[j] + mat[i][x]
				rerank(arr)
			}
		}
	}
	if len(arr) > k {
		return -1
	}
	return arr[k-1]
}
