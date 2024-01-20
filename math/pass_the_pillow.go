package math

func passThePillow(n int, time int) int {
	mod := time % (n - 1)
	div := time / (n - 1)
	if div%2 == 0 {
		return 1 + mod
	}
	return n - mod
}
