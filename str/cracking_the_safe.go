package str

import "strconv"

func crackSafe(n int, k int) string {
	power := func(a, b int) int {
		if b == 0 {
			return 1
		}
		for i := 1; i < b; i++ {
			a *= a
		}
		return a
	}
	total := power(k, n)
	ans := ""
	m := map[string]bool{}
	var dfs func(s string)
	dfs = func(s string) {
		for i := 0; i < k; i++ {
			s += strconv.Itoa(i)
			l := len(s)
			if l < n {
				dfs(s)
				s = s[:l-1]
				continue
			}
			passwd := s[l-n:]
			if _, ok := m[passwd]; ok {
				s = s[:l-1]
				continue
			}
			m[passwd] = true
			if len(m) == total {
				if ans == "" {
					ans = s
				}
				if l < len(ans) {
					ans = s
				}
				delete(m, passwd)
				s = s[:l-1]
				continue
			}
			dfs(s)
			delete(m, passwd)
			s = s[:l-1]
		}
	}
	dfs("")
	return ans
}
