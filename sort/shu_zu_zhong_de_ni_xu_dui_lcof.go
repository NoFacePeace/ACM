package sort

import "sort"

func reversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)

	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}
