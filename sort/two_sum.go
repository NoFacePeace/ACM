package sort

func twoSum(nums []int, target int) []int {
	arr := []int{}
	arr = append(arr, nums...)
	arr = quickSort(arr, 0, len(arr)-1)
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if binarySearch(arr, target-nums[i]) == -1 {
			continue
		}
		ans = append(ans, i)
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == (target - nums[i]) {
				ans = append(ans, j)
				return ans
			}
		}
		ans = []int{}
	}
	return ans
}
