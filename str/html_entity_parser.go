package str

func entityParser(text string) string {
	m := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}
	i := 0
	s := ""
	for i < len(text) {
		if text[i] != '&' {
			s += text[i : i+1]
			i++
			continue
		}
		match := false
		for j := 4; j <= 7 && i+j <= len(text); j++ {
			if _, ok := m[text[i:i+j]]; !ok {
				continue
			}
			s += m[text[i:i+j]]
			i += j
			match = true
		}
		if !match {
			s += text[i : i+1]
			i++
		}
	}
	return s
}
