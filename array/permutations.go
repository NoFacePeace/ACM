package array

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{
			nums,
		}
	}
	end := len(nums) - 1
	arr := permute(nums[:end])
	ans := [][]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= len(arr[i]); j++ {
			l := arr[i][:j]
			r := arr[i][j:]
			tmp := []int{}
			tmp = append(tmp, l...)
			tmp = append(tmp, nums[end])
			tmp = append(tmp, r...)
			ans = append(ans, tmp)
		}
	}
	return ans
}
