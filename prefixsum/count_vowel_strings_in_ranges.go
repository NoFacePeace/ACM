package prefixsum

func vowelStrings(words []string, queries [][]int) []int {
	prefix := []int{}
	n := len(words)
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i := 0; i < n; i++ {
		s := words[i]
		c1 := s[0]
		c2 := s[len(s)-1]
		cnt := 0
		if m[c1] && m[c2] {
			cnt = 1
		}
		if i == 0 {
			prefix = append(prefix, cnt)
		} else {
			prefix = append(prefix, prefix[i-1]+cnt)
		}
	}
	ans := []int{}
	for _, v := range queries {
		l, r := v[0], v[1]
		if l == 0 {
			ans = append(ans, prefix[r])
		} else {
			ans = append(ans, prefix[r]-prefix[l-1])
		}
	}
	return ans
}
