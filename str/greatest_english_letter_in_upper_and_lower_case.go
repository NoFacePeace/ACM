package str

func greatestLetter(s string) string {
	big := map[rune]bool{}
	small := map[rune]bool{}
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			small[v] = true
			continue
		}
		if v >= 'A' && v <= 'Z' {
			big[v] = true
		}
	}
	c := 'A' - 1
	for k := range big {
		if !small[k-'A'+'a'] {
			continue
		}
		if k > c {
			c = k
		}
	}
	if c == 'A'-1 {
		return ""
	}
	return string(c)
}
