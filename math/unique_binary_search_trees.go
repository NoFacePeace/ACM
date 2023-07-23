package math

func numTrees(n int) int {
	last := 1
	for i := 0; i < n; i++ {
		last = last * 2 * (2*i + 1) / (i + 2)
	}
	return last
}
