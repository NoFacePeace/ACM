package math

func countPrimes(n int) int {
	no := make([]bool, n)
	arr := []int{}
	for i := 2; i < n; i++ {
		if !no[i] {
			arr = append(arr, i)
		}
		for j := 0; j < len(arr); j++ {

			if arr[j]*i >= n {
				break
			}
			no[arr[j]*i] = true
			if i%arr[j] == 0 {
				break
			}
		}
	}
	return len(arr)
}
