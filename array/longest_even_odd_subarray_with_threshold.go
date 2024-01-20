package array

func longestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i := 0
	cnt := 0
	ans := 0
	for i < n {
		if nums[i] > threshold {
			if cnt > ans {
				ans = cnt
			}
			cnt = 0
			i++
			continue
		}
		if cnt == 0 {
			if nums[i]%2 != 0 {
				i++
			} else {
				i++
				cnt++
			}
		} else {
			if nums[i]%2 == nums[i-1]%2 {
				if cnt > ans {
					ans = cnt
				}
				if nums[i]%2 == 0 {
					cnt = 1
				} else {
					cnt = 0
				}
				i++
			} else {
				cnt++
				i++
			}
		}
	}
	if cnt > ans {
		ans = cnt
	}
	return ans
}
