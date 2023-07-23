package str

func countAsterisks(s string) int {
	cnt := 0
	ans := 0
	for _, v := range s {
		if v == '|' {
			cnt++
			continue
		}
		if v == '*' && cnt%2 == 0 {
			ans++
		}
	}
	return ans
}
