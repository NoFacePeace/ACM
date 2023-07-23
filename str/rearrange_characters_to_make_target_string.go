package str

import "math"

func rearrangeCharacters(s string, target string) int {
	m1 := map[byte]int{}
	for i := 0; i < len(target); i++ {
		m1[target[i]]++
	}
	m2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m1[s[i]]; !ok {
			continue
		}
		m2[s[i]]++
	}
	if len(m1) != len(m2) {
		return 0
	}
	ans := math.MaxInt
	for k := range m1 {
		tmp := m2[k] / m1[k]
		if tmp <= 0 {
			return 0
		}
		if tmp < ans {
			ans = tmp
		}
	}
	return ans
}
