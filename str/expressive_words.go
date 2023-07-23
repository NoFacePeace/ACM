package str

func expressiveWords(s string, words []string) int {
	cArr, nArr := count(s)
	ans := 0
	for i := 0; i < len(words); i++ {
		c, n := count(words[i])
		if check(cArr, nArr, c, n) {
			ans += 1
		}
	}
	return ans
}

func count(s string) ([]byte, []int) {
	cArr := []byte{}
	nArr := []int{}
	i := 0
	for j := 0; j < len(s); j++ {
		if j == 0 {
			cArr = append(cArr, s[j])
			nArr = append(nArr, 1)
			continue
		}
		if s[j] == cArr[i] {
			nArr[i]++
			continue
		}
		cArr = append(cArr, s[j])
		nArr = append(nArr, 1)
		i++
	}
	return cArr, nArr
}

func check(c1 []byte, n1 []int, c2 []byte, n2 []int) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
		if n1[i] == n2[i] {
			continue
		}
		if n1[i] < n2[i] {
			return false
		}
		if n1[i] == 2 {
			return false
		}
	}
	return true
}
