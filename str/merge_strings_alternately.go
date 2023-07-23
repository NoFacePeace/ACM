package str

func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	ans := ""
	i, j := 0, 0
	for i < l1 && j < l2 {
		ans += string(word1[i])
		ans += string(word2[j])
		i++
		j++
	}
	if i < l1 {
		ans += string(word1[i:])
	}
	if j < l2 {
		ans += string(word2[j:])
	}
	return ans
}
