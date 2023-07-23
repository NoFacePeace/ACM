package math

func reinitializePermutation(n int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	ans := 0
	cnt := 0
	for cnt != n {
		cnt = 0
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				tmp[i] = arr[i/2]
			} else {
				tmp[i] = arr[n/2+(i-1)/2]
			}
			if tmp[i] == i {
				cnt += 1
			}
		}
		arr = tmp
		ans += 1
	}
	return ans
}
