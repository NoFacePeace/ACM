package hash

import "strconv"

func getHint(secret string, guess string) string {
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	l := len(secret)
	for i := 0; i < l; i++ {
		m1[secret[i]]++
		m2[guess[i]]++
	}
	a := 0
	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			m1[secret[i]]--
			m2[guess[i]]--
			a++
		}
	}
	b := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for k, v := range m1 {
		b += min(v, m2[k])
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
