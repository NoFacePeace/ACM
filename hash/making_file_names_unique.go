package hash

import "strconv"

func getFolderNames(names []string) []string {
	m := map[string]int{}
	ans := []string{}
	i := 0
	for i < len(names) {
		if m[names[i]] == 0 {
			ans = append(ans, names[i])
			m[names[i]]++
			i++
			continue
		}
		cnt := m[names[i]]
		s := strconv.Itoa(cnt)
		name := names[i] + "(" + s + ")"
		if _, ok := m[name]; ok {
			m[names[i]]++
			continue
		}
		ans = append(ans, name)
		m[names[i]]++
		m[name]++
		i++
	}
	return ans
}
