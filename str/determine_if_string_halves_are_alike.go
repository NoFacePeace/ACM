package str

func halvesAreAlike(s string) bool {
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	mid := len(s) / 2
	c1 := 0
	c2 := 0
	for i := 0; i < mid; i++ {
		if _, ok := m[s[i]]; ok {
			c1++
		}
		if _, ok := m[s[mid+i]]; ok {
			c2++
		}
	}
	return c1 == c2
}
