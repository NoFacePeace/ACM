package doublepointer

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			x := n - 1
			for k := j + 1; k < n; k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				if k >= x {
					break
				}
				val := nums[i] + nums[j] + nums[k] + nums[x]
				if val == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[x]})
					continue
				}
				if val > target {
					x--
					k--
				}
			}
		}
	}
	return ans
}
