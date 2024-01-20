package str

func findTheLongestBalancedSubstring(s string) int {
	zero := 0
	one := 0
	n := len(s)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' && zero == 0 {
			one = 0
			continue
		}
		if s[i] == '1' {
			one++
			continue
		}
		if one == 0 {
			zero++
			continue
		}
		cnt := min(zero, one)
		if cnt > ans {
			ans = cnt
		}
		one = 0
		zero = 1
	}
	cnt := min(zero, one)
	if cnt > ans {
		ans = cnt
	}
	return ans * 2
}
