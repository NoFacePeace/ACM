package hash

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	m := map[int]int{}
	n := len(indices)
	for i := 0; i < n; i++ {
		src := sources[i]
		left := indices[i]
		right := len(src)
		if left+right > len(s) {
			continue
		}
		if s[left:left+right] == src {
			m[left] = i
		}
	}
	ans := ""
	i := 0
	for i < len(s) {
		if idx, ok := m[i]; ok {
			ans += targets[idx]
			i += len(sources[idx])
		} else {
			ans += s[i : i+1]
			i++
		}
	}
	return ans
}
