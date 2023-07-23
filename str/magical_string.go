package str

func magicalString(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	m := map[string]string{
		"1": "2",
		"2": "1",
	}
	s := "122"
	ans := 1
	last := "1"
	for i := 2; i < n; i++ {
		if s[i] == '1' {
			ans++
			s += last
			last = m[last]
			continue
		}
		s += last + last
		last = m[last]
	}
	return ans
}
