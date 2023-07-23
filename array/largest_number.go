package array

import (
	"strconv"
)

func largestNumber(nums []int) string {
	strs := []string{}
	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	strs = stringSort(strs)
	ans := ""
	for i := len(strs) - 1; i >= 0; i-- {
		ans += strs[i]
	}
	return ans
}

func stringSort(strs []string) []string {
	return stringQuickSort(strs, 0, len(strs)-1)
}

func stringQuickSort(strs []string, left, right int) []string {
	if left >= right {
		return strs
	}
	strs, p := stringPartiton(strs, left, right)
	strs = stringQuickSort(strs, left, p-1)
	strs = stringQuickSort(strs, p+1, right)
	return strs
}

func stringPartiton(strs []string, left, right int) ([]string, int) {
	pivot := strs[right]
	i := 0
	for j := 0; j < right; j++ {
		if isGreater(strs[j], pivot) {
			continue
		}
		strs[i], strs[j] = strs[j], strs[i]
		i++
	}
	strs[right], strs[i] = strs[i], strs[right]
	return strs, i
}

func isGreater(s1, s2 string) bool {
	i := 0
	s1, s2 = s1+s2, s2+s1
	for i < len(s1) && i < len(s2) {

		if s1[i] > s2[i] {
			return true
		}
		if s1[i] < s2[i] {
			return false
		}
		i++
	}
	return false
}
