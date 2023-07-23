package str

func alphabetBoardPath(target string) string {
	m := map[byte][2]int{}
	x, y := 0, 0
	for i := 0; i < 26; i++ {
		m[byte('a'+i)] = [2]int{x, y}
		y++
		if y >= 5 {
			y -= 5
			x++
		}
	}
	x, y = 0, 0
	c := byte('a')
	ans := ""
	for i := 0; i < len(target); i++ {
		if target[i] == c {
			ans += "!"
			continue
		}
		x1, y1 := m[target[i]][0], m[target[i]][1]
		if c == 'z' {
			for j := 0; j < x-x1; j++ {
				ans += "U"
			}
			for j := 0; j < y1-y; j++ {
				ans += "R"
			}
			c = target[i]
			x, y = x1, y1
			ans += "!"
			continue
		}
		if target[i] == 'z' {
			for j := 0; j < y-y1; j++ {
				ans += "L"
			}
			for j := 0; j < x1-x; j++ {
				ans += "D"
			}
			c = target[i]
			x, y = x1, y1
			ans += "!"
			continue
		}
		if x1 > x {
			for j := 0; j < x1-x; j++ {
				ans += "D"
			}
		} else {
			for j := 0; j < x-x1; j++ {
				ans += "U"
			}
		}
		if y1 > y {
			for j := 0; j < y1-y; j++ {
				ans += "R"
			}
		} else {
			for j := 0; j < y-y1; j++ {
				ans += "L"
			}
		}
		c = target[i]
		x, y = x1, y1
		ans += "!"
	}
	return ans
}
