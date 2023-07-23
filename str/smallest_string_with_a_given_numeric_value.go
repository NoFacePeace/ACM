package str

func getSmallestString(n int, k int) string {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	k -= n
	for i := n - 1; i >= 0; i-- {
		if k <= 25 {
			arr[i] = k + 1
			break
		}
		arr[i] = 26
		k -= 25
	}
	s := []byte{}
	for i := 0; i < n; i++ {
		s = append(s, byte(arr[i]+'a'-1))
	}
	return string(s)
}
