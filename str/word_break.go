package str

func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	l := len(wordDict)
	for i := 0; i < l; i++ {
		m[wordDict[i]] = true
	}
	var bk func(index int) bool
	bk = func(index int) bool {
		for i := index; i < len(s); i++ {
			if !m[s[index:i+1]] {
				continue
			}
			if bk(i + 1) {
				return true
			}
		}
		return false
	}
	return bk(0)
}
