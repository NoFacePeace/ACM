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

func twoSum1(nums []int, target int) []int {
	m := map[int]int{}
	ans := []int{}
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			ans = append(ans, m[target-v])
			ans = append(ans, i)
			return ans
		}
		m[v] = i
	}
	return ans
}
