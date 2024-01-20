package pointer

import "strconv"

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return nil
	}
	i := 0
	j := 0
	arr := []string{}
	for j < n {
		if i == j {
			j++
			continue
		}
		if nums[i]+j-i == nums[j] {
			j++
			continue
		}
		if nums[i] == nums[j-1] {
			s := strconv.Itoa(nums[i])
			arr = append(arr, s)
			i = j
			continue
		}
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j-1])
		arr = append(arr, si+"->"+sj)
		i = j
	}
	if nums[i] == nums[j-1] {
		s := strconv.Itoa(nums[i])
		arr = append(arr, s)
	} else {
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j-1])
		arr = append(arr, si+"->"+sj)
	}
	return arr
}
