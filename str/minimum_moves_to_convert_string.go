package str

func minimumMoves(s string) int {
	ans := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'O' {
			if cnt > 0 {
				cnt++
			}
		}
		if s[i] == 'X' {
			cnt++
		}
		if cnt == 3 {
			ans++
			cnt = 0
		}
	}
	if cnt > 0 {
		ans++
	}
	return ans
}
