package algorithm

func strStr(haystack string, needle string) int {
	kmp := func(s string) []int {
		arr := make([]int, len(s))
		for i := 0; i < len(s); i++ {
			if i == 0 {
				arr[i] = 0
				continue
			}
			p := arr[i-1]
			for s[i] != s[p] {
				if p == 0 {
					break
				}
				p = arr[p-1]
			}
			if s[i] == s[p] {
				arr[i] = p + 1
				continue
			}
			arr[i] = 0
		}
		return arr
	}
	arr := kmp(needle)
	m := 0
	i := 0
	for m < len(haystack) && i < len(needle) {
		if haystack[m] == needle[i] {
			m++
			i++
			continue
		}
		if i == 0 {
			m++
			continue
		}
		i = arr[i-1]
	}
	if i == len(needle) {
		return m - i
	}
	return -1
}
