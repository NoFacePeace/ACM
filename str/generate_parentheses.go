package str

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	left := []string{}
	right := []string{}
	for i := 0; i < n; i++ {
		left = append(left, "(")
		right = append(right, ")")
	}
	ans := []string{}
	var bk func(s string, l, r int)
	bk = func(s string, l, r int) {

		if r == n && l == n {
			ans = append(ans, s)
			return
		}
		if l == n {
			bk(s+right[r], l, r+1)
			return
		}
		if l == r {
			bk(s+left[l], l+1, r)
			return
		}
		bk(s+left[l], l+1, r)
		bk(s+right[r], l, r+1)
	}
	bk("", 0, 0)
	return ans
}
