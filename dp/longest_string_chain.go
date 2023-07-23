package dp

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(a, b int) bool {
		return len(words[a]) < len(words[b])
	})
	n := len(words)
	dp := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	match := func(s1, s2 string) bool {
		if len(s1)+1 != len(s2) {
			return false
		}
		i, j := 0, 0
		for i < len(s1) && j < len(s2) {
			if s1[i] == s2[j] {
				i++
				j++
				continue
			}
			j++
		}
		return i == len(s1)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if !match(words[j], words[i]) {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}

	return ans + 1
}
