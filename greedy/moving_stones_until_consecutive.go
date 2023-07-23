package greedy

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	left := arr[1] - arr[0] - 1
	right := arr[2] - arr[1] - 1
	if left == 0 && right == 0 {
		return []int{0, 0}
	}
	if left == 0 || right == 0 {
		return []int{1, left + right}
	}
	if left == 1 || right == 1 {
		return []int{1, left + right}
	}
	return []int{2, left + right}
}
