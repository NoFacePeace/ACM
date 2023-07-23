package array

func minOperations(nums []int, x int) int {
	l := len(nums)
	left := 0
	right := l
	sum := 0
	op := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		op += 1
	}
	total := sum
	nums = append(nums, nums...)
	ans := op
	for left <= l && right < 2*l && right-left <= l {
		if sum > x {
			if left == l {
				break
			}
			sum -= nums[left]
			left++
			op -= 1
			continue
		}
		if sum == x {
			if op < ans {
				ans = op
			}
			if left == l {
				break
			}
			sum -= nums[left]
			left++
			op -= 1
			continue
		}
		sum += nums[right]
		right++
		op += 1
	}
	if ans == l && total != x {
		return -1
	}
	return ans
}
