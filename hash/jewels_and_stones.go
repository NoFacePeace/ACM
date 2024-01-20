package hash

func numJewelsInStones(jewels string, stones string) int {
	m := map[byte]bool{}
	for _, c := range jewels {
		m[byte(c)] = true
	}
	ans := 0
	for _, c := range stones {
		if m[byte(c)] {
			ans++
		}
	}
	return ans
}
