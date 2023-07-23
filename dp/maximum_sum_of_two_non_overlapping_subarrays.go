package dp

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	l := len(nums)
	sum := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			sum[i] = nums[i]
			continue
		}
		sum[i] = nums[i] + sum[i-1]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minl := min(firstLen, secondLen)
	maxl := max(firstLen, secondLen)
	left := make([][2]int, l)
	for i := 0; i < l; i++ {
		if i < minl {
			left[i][0] = sum[i]
			left[i][1] = sum[i]
			continue
		}
		if i < maxl {
			left[i][0] = max(left[i-1][0], sum[i]-sum[i-minl])
			left[i][1] = sum[i]
			continue
		}
		left[i][0] = max(left[i-1][0], sum[i]-sum[i-minl])
		left[i][1] = max(left[i-1][1], sum[i]-sum[i-maxl])
	}
	right := make([][2]int, l)
	for i := l - 1; i >= 0; i-- {
		if l-1-i <= minl {
			right[i][0] = sum[l-1] - sum[i]
			right[i][1] = sum[l-1] - sum[i]
			continue
		}
		if l-1-i <= maxl {
			right[i][0] = max(right[i+1][0], sum[i+minl]-sum[i])
			right[i][1] = sum[l-1] - sum[i]
			continue
		}
		right[i][0] = max(right[i+1][0], sum[i+minl]-sum[i])
		right[i][1] = max(right[i+1][1], sum[i+maxl]-sum[i])
	}
	ans := 0
	for i := 0; i < l; i++ {
		ans = max(ans, right[i][1]+left[i][0])
		ans = max(ans, right[i][0]+left[i][1])
	}
	return ans
}
