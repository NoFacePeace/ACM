package array

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	ans := make([]int, k)
	selectSort := func(end int) {
		for i := end; i > 0; i-- {
			if m[ans[i]] < m[ans[i-1]] {
				return
			}
			ans[i], ans[i-1] = ans[i-1], ans[i]
		}
	}
	i := 0
	for j, v := range m {
		if i == 0 {
			ans[i] = j
			i++
			continue
		}
		if i < k {
			ans[i] = j
			selectSort(i)
			i++
			continue
		}
		if m[ans[i-1]] > v {
			continue
		}
		ans[i-1] = j
		selectSort(i - 1)
	}
	return ans
}
