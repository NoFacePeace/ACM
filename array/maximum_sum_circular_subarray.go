package array

func maxSubarraySumCircular(nums []int) int {
	nums = append(nums, nums...)
	sum := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum = append(sum, nums[i])
			continue
		}
		sum = append(sum, nums[i]+sum[i-1])
	}
	l := len(sum)
	if l == 0 {
		return 0
	}
	ans := sum[0]
	q := []int{0}
	for i := 1; i < l; i++ {
		if i-q[0] > l/2 {
			q = q[1:]
		}
		if sum[i]-sum[q[0]] > ans {
			ans = sum[i] - sum[q[0]]
		}
		for len(q) > 0 && sum[q[len(q)-1]] >= sum[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return ans
}
