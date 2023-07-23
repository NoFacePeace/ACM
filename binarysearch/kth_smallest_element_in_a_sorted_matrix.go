package binarysearch

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	left := matrix[0][0]
	right := matrix[n-1][n-1]
	check := func(mid int) bool {
		i := n - 1
		j := 0
		cnt := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				cnt += i + 1
				j++
				continue
			}
			i--
		}
		return cnt >= k
	}
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
			continue
		}
		left = mid + 1
	}
	return left
}
