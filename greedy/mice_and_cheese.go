package greedy

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	if n < k {
		return 0
	}
	dst := []int{}
	sum := 0
	for i := 0; i < n; i++ {
		dst = append(dst, reward2[i]-reward1[i])
		sum += reward2[i]
	}
	sort.Ints(dst)
	for i := 0; i < k; i++ {
		sum -= dst[i]
	}
	return sum
}
