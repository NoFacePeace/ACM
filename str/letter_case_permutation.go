package str

func letterCasePermutation(s string) []string {
	var bk func(s string, n string, i int)
	ans := []string{}
	bk = func(s string, n string, i int) {
		if len(s) == i {
			ans = append(ans, n)
			return
		}
		if s[i] >= '0' && s[i] <= '9' {
			tmp := string(s[i])
			bk(s, n+tmp, i+1)
			return
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			tmp := string(s[i])
			bk(s, n+tmp, i+1)
			tmp = string('A' + s[i] - 'a')
			bk(s, n+tmp, i+1)
			return
		}
		tmp := string(s[i])
		bk(s, n+tmp, i+1)
		tmp = string('a' + s[i] - 'A')
		bk(s, n+tmp, i+1)
	}
	bk(s, "", 0)
	return ans
}
