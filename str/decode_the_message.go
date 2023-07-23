package str

func decodeMessage(key string, message string) string {
	m := map[byte]int{}
	cnt := 0
	for i := 0; i < len(key); i++ {
		if key[i] == ' ' {
			continue
		}
		if _, ok := m[key[i]]; ok {
			continue
		}
		m[key[i]] = 'a' + cnt
		cnt++
	}
	ans := ""
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			ans += " "
			continue
		}
		ans += string(byte(m[message[i]]))
	}
	return ans
}
