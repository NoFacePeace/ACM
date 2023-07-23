package tree

import "sort"

func reversePairs(nums []int) int {
	tmp := []int{}
	for _, v := range nums {
		tmp = append(tmp, v)
	}
	sort.Ints(tmp)
	m := map[int]int{}
	cnt := 0
	for _, v := range tmp {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = cnt
		cnt++
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = m[nums[i]] + 1
	}
	tree := make([]int, n+1)
	lowbit := func(x int) int {
		return x & (-x)
	}
	query := func(x int) int {
		cnt := 0
		for x > 0 {
			cnt += tree[x]
			x -= lowbit(x)
		}
		return cnt
	}
	update := func(x int) {
		for x < n {
			tree[x] += 1
			x += lowbit(x)
		}
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += query(nums[i] - 1)
		update(nums[i])
	}
	return ans
}
