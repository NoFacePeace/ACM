package array

import "sort"

func countDifferentSubsequenceGCDs(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	max := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	m := map[int]bool{}
	inc := 0
	for inc < max {
		inc++
		arr := []int{}
		for i := 0; i < len(nums); i++ {
			if nums[i]%inc == 0 {
				arr = append(arr, nums[i])
			}
		}
		if len(arr) == 0 {
			continue
		}
		if len(arr) == 1 {
			if arr[0] == inc {
				m[inc] = true
			}
			continue
		}
		d := arr[0]
		for i := 1; i < len(arr); i++ {
			d = gcd(d, arr[i])
			if d == inc {
				m[inc] = true
				break
			}
		}
	}
	return len(m)
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
