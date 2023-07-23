package dp

func longestDecomposition(text string) int {
	var f func(left, right int) int
	dp := make([][]int, len(text))
	for i := 0; i < len(text); i++ {
		dp[i] = make([]int, len(text))
	}
	f = func(left, right int) int {
		if left > right {
			return 0
		}
		if left == right {
			return 1
		}
		if dp[left][right] != 0 {
			return dp[left][right]
		}
		max := 1
		mid := (left + right) / 2
		for i := left; i <= mid; i++ {
			s1 := text[left : i+1]
			s2 := text[right-i+left : right+1]
			if s1 != s2 {
				continue
			}
			val := f(i+1, right-i+left-1)
			if val+2 > max {
				max = val + 2
			}
		}
		dp[left][right] = max
		return max
	}
	return f(0, len(text)-1)
}
