package str

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	m := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	ans := 0
	for i := 0; i < len(items); i++ {
		if items[i][m[ruleKey]] == ruleValue {
			ans++
		}
	}
	return ans
}
