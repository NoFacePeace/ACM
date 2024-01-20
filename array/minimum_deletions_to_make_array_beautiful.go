package array

func minDeletion(nums []int) int {
	n := len(nums)
	cnt := 0
	i := 0
	for i < n {
		if i == n-1 {
			cnt++
			break
		}
		if nums[i] != nums[i+1] {
			i += 2
			continue
		}
		cnt++
		i++
	}
	return cnt
}
