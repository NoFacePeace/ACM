package math

func getRow(rowIndex int) []int {
	ans := []int{1}
	for i := 0; i < rowIndex; i++ {
		last := ans
		ans = make([]int, len(last)+1)
		for j := 0; j < len(ans); j++ {
			if j == 0 {
				ans[j] = last[j]
				continue
			}
			if j == len(ans)-1 {
				ans[j] = last[j-1]
				continue
			}
			ans[j] = last[j] + last[j-1]
		}
	}
	return ans
}
