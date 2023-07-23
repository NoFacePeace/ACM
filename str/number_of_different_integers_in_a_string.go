package str

func numDifferentIntegers(word string) int {
	str := ""
	m := map[string]bool{}
	for _, v := range word {
		if v == '0' {
			if str == "0" {
				continue
			}
			str += string(v)
			continue
		}
		if v >= '1' && v <= '9' {
			if str == "0" {
				str = string(v)
				continue
			}
			str += string(v)
			continue
		}
		if str == "" {
			continue
		}
		m[str] = true
		str = ""
	}
	if str != "" {
		m[str] = true
	}
	return len(m)
}
