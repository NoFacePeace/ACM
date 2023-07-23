package math

import "sort"

func maxScore(nums []int) int {
	l := len(nums)
	res := make([][]int, l)
	for i := 0; i < l; i++ {
		res[i] = make([]int, l)
	}
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			g := gcd(nums[i], nums[j])
			res[i][j] = g
		}
	}
	ans := 0
	var bk func(sum, index int)
	used := make([]bool, l)

	bk = func(sum, index int) {
		if index == 0 {
			if sum > ans {
				ans = sum
			}
			return
		}
		for i := 0; i < l; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			for j := i + 1; j < l; j++ {
				if used[j] {
					continue
				}
				used[j] = true
				bk(sum+index*res[i][j], index-1)
				used[j] = false
			}
			used[i] = false
		}
	}
	bk(0, l/2)
	return ans
}
