package dp

func countSubstrings(s string, t string) int {
	l1 := len(s)
	l2 := len(t)
	left := make([][]int, l1+1)
	right := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		left[i] = make([]int, l2+1)
		right[i] = make([]int, l2+1)
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if s[i] != t[j] {
				continue
			}
			left[i+1][j+1] = left[i][j] + 1
		}
	}
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if s[i] != t[j] {
				continue
			}
			right[i][j] = right[i+1][j+1] + 1
		}
	}
	ans := 0
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if s[i] == t[j] {
				continue
			}
			ans += (left[i][j] + 1) * (right[i+1][j+1] + 1)
		}
	}
	return ans
}
