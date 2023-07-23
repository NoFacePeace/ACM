package greedy

func prevPermOpt1(arr []int) []int {
	l := len(arr)
	for i := l - 2; i >= 0; i-- {
		if arr[i] <= arr[i+1] {
			continue
		}
		for j := l - 1; j > i; j-- {
			if arr[j] >= arr[i] {
				continue
			}
			if arr[j] == arr[j-1] {
				continue
			}
			arr[i], arr[j] = arr[j], arr[i]
			return arr
		}
	}
	return arr
}
