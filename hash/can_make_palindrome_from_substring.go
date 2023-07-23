package hash

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	check := func(left, right, k int) bool {
		if arr[left][right] != 0 {
			return arr[left][right] <= k
		}
		m := map[byte]int{}
		ss := s[left : right+1]
		for i := 0; i < len(ss); i++ {
			c := s[i]
			m[c]++
			m[c] = m[c] % 2
		}
		cnt := 0
		for _, v := range m {
			cnt += v
		}
		arr[left][right] = cnt / 2
		return cnt/2 <= k
	}
	ans := []bool{}
	for _, v := range queries {
		left, right, k := v[0], v[1], v[2]
		ok := check(left, right, k)
		ans = append(ans, ok)
	}
	return ans
}
