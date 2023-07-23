package array

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ans := 0
	if sum == goal {
		return ans
	}
	if sum > goal {
		sum, goal = goal, sum
	}
	for sum != goal {
		diff := goal - sum
		if diff >= limit {
			sum += diff / limit * limit
			ans += diff / limit
			continue
		}
		sum += diff
		ans++
	}
	return ans
}
