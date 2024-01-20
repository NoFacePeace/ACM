package hash

func findRepeatedDnaSequences(s string) []string {
	hash := map[string]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		if i+10 > n {
			break
		}
		hash[s[i:i+10]]++
	}
	ans := []string{}
	for k, v := range hash {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
