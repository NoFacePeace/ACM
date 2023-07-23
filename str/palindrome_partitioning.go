package str

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{
			{s},
		}
	}
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if j+i >= l {
				break
			}
			if i == 0 {
				dp[j][j+i] = true
				continue
			}
			if i == 1 {
				if s[j] == s[j+i] {
					dp[j][j+i] = true
					continue
				}
				dp[j][j+i] = false
				continue
			}
			if s[j] == s[j+i] && dp[j+1][j+i-1] {
				dp[j][j+i] = true
				continue
			}
			dp[j][j+i] = false
		}
	}
	ans := [][]string{}
	split := []string{}
	var bk func(i, j int)
	bk = func(i, j int) {
		if !dp[i][j] {
			return
		}
		if j == l-1 {
			tmp := []string{}
			tmp = append(tmp, split...)
			tmp = append(tmp, s[i:j+1])
			ans = append(ans, tmp)
			return
		}
		for k := j + 1; k < l; k++ {
			split = append(split, s[i:j+1])
			bk(j+1, k)
			split = split[:len(split)-1]
		}
	}
	for i := 0; i < l; i++ {
		bk(0, i)
	}
	return ans
}
