package str

func minimumDeletions(s string) int {
	l := len(s)
	a := make([]int, l)
	for i := 0; i < l; i++ {
		cnt := 0
		if s[i] == 'b' {
			cnt += 1
		}
		if i != 0 {
			cnt += a[i-1]
		}
		a[i] = cnt
	}
	b := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		cnt := 0
		if s[i] == 'a' {
			cnt += 1
		}
		if i != l-1 {
			cnt += b[i+1]
		}
		b[i] = cnt
	}
	ans := l
	for i := 0; i < l; i++ {
		if a[i]+b[i]-1 < ans {
			ans = a[i] + b[i] - 1
		}
	}
	return ans
}
