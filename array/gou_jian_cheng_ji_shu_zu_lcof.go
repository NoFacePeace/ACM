package array

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []int{0}
	}
	left := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = a[i]
			continue
		}
		left[i] = a[i] * left[i-1]
	}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = a[i]
			continue
		}
		right[i] = a[i] * right[i+1]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ans[i] = right[i+1]
			continue
		}
		if i == n-1 {
			ans[i] = left[i-1]
			continue
		}
		ans[i] = left[i-1] * right[i+1]
	}
	return ans
}
