package str

func countSeniors(details []string) int {
	age := 11
	cnt := 0
	for _, s := range details {
		if s[age] == '6' && s[age+1] > '0' {
			cnt++
		}
		if s[age] > '6' {
			cnt++
		}
	}
	return cnt
}
