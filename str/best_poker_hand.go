package str

func bestHand(ranks []int, suits []byte) string {
	m1 := map[int]int{}
	m2 := map[byte]int{}
	for i := 0; i < len(ranks); i++ {
		m1[ranks[i]]++
		m2[suits[i]]++
	}
	for _, v := range m2 {
		if v == 5 {
			return "Flush"
		}
	}
	ans := "High Card"
	for _, v := range m1 {
		if v >= 3 {
			return "Three of a Kind"
		}
		if v == 2 {
			ans = "Pair"
		}
	}
	return ans
}
