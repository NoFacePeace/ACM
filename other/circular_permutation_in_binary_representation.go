package other

func circularPermutation(n int, start int) []int {
	max := 1
	for i := 0; i < n; i++ {
		max = max * 2
	}
	m := map[string]bool{}
	var dfs func(s string) bool
	arr := []string{}
	dfs = func(s string) bool {
		if len(arr) == max {
			cnt := 0
			for i := 0; i < len(arr[0]); i++ {
				if arr[0][i] != arr[max-1][i] {
					cnt++
				}
			}
			return cnt == 1
		}
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				tmp := s[:i] + "1" + s[i+1:]
				if m[tmp] {
					continue
				}
				m[tmp] = true
				arr = append(arr, tmp)
				if dfs(tmp) {
					return true
				}
				m[tmp] = false
				arr = arr[:len(arr)-1]
			} else {
				tmp := s[:i] + "0" + s[i+1:]
				if m[tmp] {
					continue
				}
				m[tmp] = true
				arr = append(arr, tmp)
				if dfs(tmp) {
					return true
				}
				m[tmp] = false
				arr = arr[:len(arr)-1]
			}
		}
		return false
	}
	num := start
	s := ""
	for num != 0 {
		if num&1 == 1 {
			s = "1" + s
		} else {
			s = "0" + s
		}
		num = num >> 1
	}
	for len(s) < n {
		s = "0" + s
	}
	arr = append(arr, s)
	ans := []int{}
	m[s] = true
	dfs(s)
	for i := 0; i < len(arr); i++ {
		num := 0
		p := 1
		for j := len(arr[i]) - 1; j >= 0; j-- {
			if arr[i][j] == '1' {
				num += p
			}
			p = p * 2
		}
		ans = append(ans, num)
	}
	return ans
}
