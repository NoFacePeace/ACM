package backtrack

func punishmentNumber(n int) int {
	ans := 0
	satisfy := func(mult, num int) bool {
		arr := []int{}
		for mult != 0 {
			arr = append(arr, mult%10)
			mult /= 10
		}
		var dfs func(sum, bit, i int) bool
		dfs = func(sum, bit, i int) bool {
			if sum > num {
				return false
			}
			if i == len(arr) {
				if sum == num {
					return true
				}
				return false
			}
			if dfs(sum+arr[i], 10, i+1) {
				return true
			}
			if dfs(sum+arr[i]*bit, bit*10, i+1) {
				return true
			}
			return false
		}
		return dfs(0, 1, 0)
	}
	for i := 1; i <= n; i++ {
		if satisfy(i*i, i) {
			ans += i * i
		}
	}
	return ans
}
