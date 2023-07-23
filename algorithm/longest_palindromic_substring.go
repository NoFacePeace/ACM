package algorithm

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	tmp := ""
	for i := 0; i < len(s); i++ {
		tmp += "#" + string(s[i])
	}
	tmp += "#"
	s = tmp
	j, right := -1, -1
	arr := make([]int, len(s))
	expand := func(left, right int) int {
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			left--
			right++
		}
		return (right - left - 2) / 2
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	start := 0
	end := -1
	for i := 0; i < len(s); i++ {
		if right < i {
			arr[i] = expand(i, i)
		} else {
			l := min(arr[2*j-i], right-i)
			arr[i] = expand(i-l, i+l)
		}
		if i+arr[i] > right {
			right = i + arr[i]
			j = i
		}
		if arr[i]*2+1 > end-start {
			start = i - arr[i]
			end = i + arr[i]
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] == '#' {
			continue
		}
		ans += string(s[i])
	}
	return ans
}
