package str

func balancedString(s string) int {
	l := len(s)
	m := make([]map[byte]int, l)
	for i := 0; i < l; i++ {
		m[i] = map[byte]int{}
		if i == 0 {
			m[i][s[i]]++
			continue
		}
		m[i]['Q'] = m[i-1]['Q']
		m[i]['W'] = m[i-1]['W']
		m[i]['E'] = m[i-1]['E']
		m[i]['R'] = m[i-1]['R']
		m[i][s[i]] = m[i-1][s[i]] + 1
	}
	codition := map[byte]int{}
	for k, v := range m[l-1] {
		if v-l/4 > 0 {
			codition[k] = v - l/4
		}
	}
	if len(codition) <= 0 {
		return 0
	}
	le := -1
	ri := 0
	ans := l
	satisfy := func(left, right int) bool {
		if left == -1 {
			for k, v := range codition {
				if m[right][k] < v {
					return false
				}
			}
			if right+1 < ans {
				ans = right + 1
			}
			return true
		}
		for k, v := range codition {
			if m[right][k]-m[left][k] < v {
				return false
			}
		}
		if right-left < ans {
			ans = right - left
		}
		return true
	}
	for ri < l {
		if satisfy(le, ri) {
			le++
			continue
		}
		ri++
	}
	return ans
}
