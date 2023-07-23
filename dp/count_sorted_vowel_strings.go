package dp

func countVowelStrings(n int) int {
	ans := make([]int, 5)
	i := 1
	for i <= n {
		for j := 4; j >= 0; j-- {
			if i == 1 {
				ans[j] = 1
				continue
			}
			if j == 4 {
				continue
			}
			ans[j] += ans[j+1]
		}
		i++
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum += ans[i]
	}
	return sum
}
