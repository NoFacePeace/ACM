package hash

func vowelStrings(words []string, left int, right int) int {
	cnt := 0
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i := left; i <= right; i++ {
		s := words[i]
		n := len(s)
		if n == 0 {
			continue
		}
		if !m[s[0]] {
			continue
		}
		if !m[s[n-1]] {
			continue
		}
		cnt++
	}
	return cnt
}
