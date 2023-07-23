package str

func countTime(time string) int {
	m1 := 1
	m2 := 1
	if time[4] == '?' {
		m2 = 10
	}
	if time[3] == '?' {
		m1 = 6
	}
	if time[1] == '?' && time[0] == '?' {
		return 24 * m2 * m1
	}
	if time[1] == '?' && time[0] < '2' {
		return 10 * m2 * m1
	}
	if time[1] == '?' && time[0] == '2' {
		return 4 * m2 * m1
	}
	if time[0] == '?' && time[1] < '4' {
		return 3 * m2 * m1
	}
	if time[0] == '?' {
		return 2 * m2 * m1
	}
	return m1 * m2
}
