package array

func partitionDisjoint(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	now := nums[0]
	max := nums[0]
	for i := 0; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
		if nums[i] > now {
			continue
		}
		if nums[i] == now {
			continue
		}
		if now == max {
			continue
		}
		now = max
	}
	repeat := -1
	ans := 0
	for i := 0; i < l; i++ {
		if nums[i] > now {
			break
		}
		if nums[i] == now {
			ans++
			repeat++
			continue
		}
		ans++
		if repeat == -1 {
			continue
		}
		repeat = 0
	}
	return ans - repeat
}
