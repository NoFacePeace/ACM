package array

func threeSum(nums []int) [][]int {
	nums = qSort(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k > j && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

func qSort(nums []int) []int {
	partition := func(nums []int, left, right int) ([]int, int) {
		pivot := nums[right]
		i := left
		for j := left; j < right; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		return nums, i
	}
	var sort func(num []int, left, right int) []int
	sort = func(nums []int, left, right int) []int {
		if left >= right {
			return nums
		}
		nums, p := partition(nums, left, right)
		nums = sort(nums, left, p-1)
		nums = sort(nums, p+1, right)
		return nums
	}
	return sort(nums, 0, len(nums)-1)
}
