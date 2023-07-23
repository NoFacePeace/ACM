package str

func oddString(words []string) string {
	n := len(words)
	if n == 0 {
		return ""
	}
	arr := []byte{}
	for i := 1; i < len(words[0]); i++ {
		arr = append(arr, words[0][i]-words[0][i-1])
	}
	diff := 0
	idx := 0
	for i := 1; i < n; i++ {
		for j := 1; j < len(words[i]); j++ {
			if words[i][j]-words[i][j-1] != arr[j-1] {
				idx = i
				diff++
				break
			}
		}
	}
	if diff > 1 {
		return words[0]
	}
	return words[idx]
}
