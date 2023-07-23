package dp

func shortestCommonSupersequence(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	dp := make([][]string, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]string, l2+1)
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + string(str1[i])
				continue
			}
			if len(dp[i][j+1]) > len(dp[i+1][j]) {
				dp[i+1][j+1] = dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	sub := dp[l1][l2]
	i := 0
	j := 0
	l3 := len(sub)
	k := 0
	ans := ""
	for i < l1 && j < l2 && k < l3 {
		if str1[i] == sub[k] && str2[j] == sub[k] {
			ans += string(sub[k])
			k++
			i++
			j++
			continue
		}
		if str1[i] != sub[k] {
			ans += string(str1[i])
			i++
		}
		if str2[j] != sub[k] {
			ans += string(str2[j])
			j++
		}
	}
	if i < l1 {
		ans += str1[i:]
	}
	if j < l2 {
		ans += str2[j:]
	}
	return ans
}
