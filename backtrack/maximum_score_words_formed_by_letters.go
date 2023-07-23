package backtrack

func maxScoreWords(words []string, letters []byte, score []int) int {
	max := 0
	m := map[byte]int{}
	for _, v := range letters {
		m[v]++
	}
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if i == len(words) {
			if s > max {
				max = s
				return
			}
		}
		dfs(i+1, s)
		j := 0
		for ; j < len(words[i]); j++ {
			if m[words[i][j]]-1 < 0 {
				break
			}
			m[words[i][j]]--
			s += score[words[i][j]-'a']
		}
		if j == len(words[i]) {
			dfs(i+1, s)
		}
		j--
		for ; j >= 0; j-- {
			m[words[i][j]]++
		}
	}
	dfs(0, 0)
	return max
}
