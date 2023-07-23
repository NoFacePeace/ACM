package dp

func maximumSum(arr []int) int {
	n := len(arr)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = arr[i]
			continue
		}
		if left[i-1] < 0 {
			left[i] = arr[i]
			continue
		}
		left[i] = left[i-1] + arr[i]
	}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = arr[i]
			continue
		}
		if right[i+1] < 0 {
			right[i] = arr[i]
			continue
		}
		right[i] = right[i+1] + arr[i]
	}
	max := left[n-1]
	for i := 0; i < n; i++ {
		if left[i] == arr[i] && right[i] == arr[i] {
			if arr[i] > max {
				max = arr[i]
			}
			continue
		}
		if left[i]-arr[i]+right[i]-arr[i] > max {
			max = left[i] - arr[i] + right[i] - arr[i]
		}
	}
	return max
}
