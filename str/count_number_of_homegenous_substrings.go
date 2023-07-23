package str

func countHomogenous(s string) int {
	same := 1
	ans := 0
	m := map[int]int{}
	mod := int(1e9 + 7)
	compute := func(num int) int {
		if _, ok := m[num]; ok {
			return m[num]
		}
		cnt := 0
		for i := 1; i <= num; i++ {
			cnt += i
		}
		m[num] = cnt % mod
		return m[num]
	}
	for i := 1; i <= len(s); i++ {
		if i == len(s) {
			ans += compute(same)
			continue
		}
		if s[i] != s[i-1] {
			ans += compute(same)
			same = 1
			continue
		}
		same++
	}
	return ans
}
