package doublepointer

func lastSubstring(s string) string {
	i, j := 0, 1
	n := len(s)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j = j + k + 1
		}
	}
	return s[i:]
}
