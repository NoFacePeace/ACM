package str

func largestMerge(word1 string, word2 string) string {
	i := 0
	j := 0
	s := ""
	d := 0
	for i < len(word1) && j < len(word2) {
		if word1[i:] == word2[j:] {
			s += string(word1[i])
			i++
			continue
		}
		if word1[i] > word2[j] {
			i -= d
			s += string(word1[i])
			i++
			j -= d
			d = 0
			continue
		}
		if word1[i] < word2[j] {
			j -= d
			s += string(word2[j])
			j++
			i -= d
			d = 0
			continue
		}
		i++
		j++
		d++
	}
	if d > 0 {
		i -= d
		j -= d
	}
	s1 := ""
	for i < len(word1) {
		s1 += string(word1[i])
		i++
	}
	s2 := ""
	for j < len(word2) {
		s2 += string(word2[j])
		j++
	}
	if s+s1+s2 > s+s2+s1 {
		return s + s1 + s2
	}
	return s + s2 + s1
}
