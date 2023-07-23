package str

func groupAnagrams(strs []string) [][]string {
	index := 0
	m := map[string]int{}
	ans := [][]string{}
	for i := 0; i < len(strs); i++ {
		s := stringSort(strs[i])
		v, ok := m[s]
		if !ok {
			tmp := []string{}
			tmp = append(tmp, strs[i])
			ans = append(ans, tmp)
			m[s] = index
			index++
			continue
		}
		ans[v] = append(ans[v], strs[i])
	}
	return ans
}

func stringSort(s string) string {
	if len(s) <= 1 {
		return s
	}
	partition := func(s string) (string, int) {
		arr := []byte(s)
		pivot := arr[len(arr)-1]
		i := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > pivot {
				continue
			}
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		arr[i], arr[len(arr)-1] = pivot, arr[i]
		return string(arr), i
	}
	s, p := partition(s)
	s1 := ""
	s2 := ""
	if p > 0 {
		s1 = stringSort(s[:p])
	}
	if p < len(s)-1 {
		s2 = stringSort(s[p+1:])
	}
	return s1 + s[p:p+1] + s2
}
