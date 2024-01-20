package math

func sumNums(n int) int {
	ans := 0
	var f func(n int) bool
	f = func(n int) bool {
		ans += n
		return n > 0 && f(n-1)
	}
	f(n)
	return ans
}
