package str

func customSortString(order string, s string) string {
	m := map[int]int{}
	for _, v := range order {
		index := int(v - 'a')
		m[index] = 0
	}
	suffix := ""
	for _, v := range s {
		index := int(v - 'a')
		if _, ok := m[index]; ok {
			m[index]++
			continue
		}
		suffix += string(v)
	}
	ans := ""
	for _, v := range order {
		index := int(v - 'a')
		for i := 0; i < m[index]; i++ {
			ans += string(v)
		}
	}
	return ans + suffix
}
