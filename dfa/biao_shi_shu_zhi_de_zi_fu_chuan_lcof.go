package dfa

func isNumber(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	m := map[byte]bool{
		'+': true,
		'0': true,
		'.': true,
		' ': true,
	}
	e := 0
	dot := 0
	num := 0
	for _, v := range s {
		if v == ' ' {
			if !m[' '] {
				return false
			}
			if num != 0 {
				m = map[byte]bool{
					' ': true,
					'f': true,
				}
			}
			continue
		}
		if v == 'e' || v == 'E' {
			if !m['e'] {
				return false
			}
			e++
			m = map[byte]bool{
				'+': true,
				'0': true,
			}
			continue
		}
		if v == '.' {
			if !m['.'] {
				return false
			}
			m = map[byte]bool{
				'0': true,
			}
			if num != 0 {
				m['f'] = true
				m[' '] = true
			}
			if num != 0 && e == 0 {
				m['e'] = true
			}
			dot++
			continue
		}
		if v == '+' || v == '-' {
			if !m['+'] {
				return false
			}
			m = map[byte]bool{
				'0': true,
			}
			if dot == 0 {
				m['.'] = true
			}
			continue
		}
		if v >= '0' && v <= '9' {
			if !m['0'] {
				return false
			}
			m = map[byte]bool{
				'0': true,
				'f': true,
				' ': true,
			}
			if e == 0 {
				m['e'] = true
			}
			if e == 0 && dot == 0 {
				m['.'] = true
			}
			num++
			continue
		}
		return false
	}
	if m['f'] {
		return true
	}
	return false
}
