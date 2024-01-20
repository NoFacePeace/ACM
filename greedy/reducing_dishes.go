package greedy

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	ans := 0
	cnt := 0
	n := len(satisfaction)
	for i := n - 1; i >= 0; i-- {
		if satisfaction[i]+cnt < 0 {
			break
		}
		ans += cnt + satisfaction[i]
		cnt += satisfaction[i]
	}
	return ans
}
