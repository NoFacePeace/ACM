package str

func checkIfPangram(sentence string) bool {
	m := map[byte]bool{}
	for i := 0; i < len(sentence); i++ {
		m[sentence[i]] = true
	}
	return len(m) == 26
}
