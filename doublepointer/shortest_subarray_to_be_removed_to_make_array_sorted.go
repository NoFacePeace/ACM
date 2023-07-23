package doublepointer

func findLengthOfShortestSubarray(arr []int) int {
	l := len(arr)
	if l == 0 {
		return l
	}
	left := 0
	for i := 1; i < l; i++ {
		if arr[i] < arr[i-1] {
			break
		}
		left++
	}
	right := l - 1
	for i := l - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			break
		}
		right--
	}
	if right <= left {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := max(left+1, l-right)
	for i := right; i < l; i++ {
		for j := left; j >= 0; j-- {
			if arr[i] >= arr[j] {
				ans = max(ans, j+l-i+1)
				break
			}
		}
	}
	return l - ans
}
