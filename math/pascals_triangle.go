package math

func generate(numRows int) [][]int {
	ans := [][]int{}
	for i := 0; i < numRows; i++ {
		if i == 0 {
			ans = append(ans, []int{1})
			continue
		}
		ans = append(ans, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 {
				ans[i] = append(ans[i], ans[i-1][j])
				continue
			}
			if j == i {
				ans[i] = append(ans[i], ans[i-1][j-1])
				continue
			}
			ans[i] = append(ans[i], ans[i-1][j]+ans[i-1][j-1])
		}
	}
	return ans
}
