package array

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{
			{},
		}
	}
	arr := subsets(nums[:len(nums)-1])
	ans := [][]int{}
	ans = append(ans, arr...)
	for i := 0; i < len(arr); i++ {
		tmp := []int{}
		tmp = append(tmp, arr[i]...)
		tmp = append(tmp, nums[len(nums)-1])
		ans = append(ans, tmp)
	}

	return ans
}
