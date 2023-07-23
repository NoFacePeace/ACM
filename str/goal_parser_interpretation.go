package str

func interpret(command string) string {
	ans := ""
	i := 0
	l := len(command)
	for i < l {
		if command[i] == 'G' {
			ans += "G"
			i++
			continue
		}
		if command[i] == '(' && i+1 < l && command[i+1] == ')' {
			ans += "o"
			i += 2
			continue
		}
		ans += "al"
		i += 4
	}
	return ans
}
