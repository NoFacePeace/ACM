package array

func dieSimulator(n int, rollMax []int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 6)
	}
	sum := func(arr []int, j int) int {
		s := 0
		for i := 0; i < len(arr); i++ {
			if i == j {
				continue
			}
			s += arr[i]
		}
		return s
	}
	mod := int(1e9) + 7
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			if i == 0 {
				arr[i][j] = 1
				continue
			}
			if i-rollMax[j] < 0 {
				arr[i][j] = sum(arr[i-1], -1) % mod
				continue
			}
			for k := 0; k < rollMax[j]; k++ {
				arr[i][j] += sum(arr[i-k-1], j) % mod
			}
		}
	}
	return sum(arr[n-1], -1) % mod
}
