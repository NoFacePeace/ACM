package str

func minOperations(s string) int {
	c1 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '1' {
				c1++
			}
		} else {
			if s[i] == '0' {
				c1++
			}
		}
	}
	c2 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				c2++
			}
		} else {
			if s[i] == '1' {
				c2++
			}
		}
	}
	if c1 < c2 {
		return c1
	}
	return c2
}
