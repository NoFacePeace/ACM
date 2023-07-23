package greedy

func longestPalindrome(words []string) int {
	m := map[string]int{}
	ans := 0
	for _, v := range words {
		l := len(v)
		s := ""
		mid := l / 2
		if l%2 == 0 {
			s = v[mid:l] + v[0:mid]
		} else {
			s = v[mid+1:] + string(v[mid]) + v[0:mid]
		}
		if m[s] <= 0 {
			m[v]++
			continue
		}
		ans += l * 2
		m[s]--
	}
	max := 0
	for k, v := range m {
		if v == 0 {
			continue
		}
		l := len(k)
		mid := l / 2
		s := ""
		if l%2 == 0 {
			s = k[mid:l] + k[0:mid]
		} else {
			s = k[mid+1:] + string(k[mid]) + k[0:mid]
		}
		if s != k {
			continue
		}
		if l > max {
			max = l
		}
	}
	return ans + max
}
