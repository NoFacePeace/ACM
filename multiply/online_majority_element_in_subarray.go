package multiply

import "sort"

type MajorityChecker struct {
	segT  []*SegTree
	count map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	count := make(map[int][]int)
	// 统计每个数字出现的索引，并分成单独的数组
	for i := 0; i < len(arr); i++ {
		if _, ok := count[arr[i]]; !ok {
			count[arr[i]] = []int{}
		}
		count[arr[i]] = append(count[arr[i]], i)
	}
	return MajorityChecker{
		segT:  buildSegTree(arr),
		count: count,
	}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	ret := &SegTree{}
	// 先通过线段树找到需要计算的数字
	ret = querySegTree(this.segT, 1, 0, len(this.segT)/2-1, left, right, ret)
	if ret == nil {
		return -1
	}
	c, v := this.count, ret.val
	if _, ok := c[v]; !ok {
		return -1
	}

	// 二分查找找到左右点对应的索引
	start := sort.Search(len(c[v]), func(i int) bool { return left <= c[v][i] })
	end := sort.Search(len(c[v]), func(i int) bool { return right < c[v][i] })
	if (end - start) >= threshold {
		return v
	}
	return -1
}

type SegTree struct {
	val, count int
}

func buildSegTree(arr []int) []*SegTree {
	n := len(arr)
	depth := 1
	for depth < n {
		depth *= 2
	}
	segT := make([]*SegTree, depth*2)
	for i := 0; i < n; i++ {
		segT[depth+i] = &SegTree{arr[i], 1}
	}
	for i := depth - 1; i > 0; i-- {
		segT[i] = combineSegTreeChild(segT[i*2], segT[i*2+1])
	}
	return segT
}

func combineSegTreeChild(left, right *SegTree) *SegTree {
	if left == nil {
		left = &SegTree{}
	}
	if right == nil {
		right = &SegTree{}
	}
	var val, count int
	if left.val == right.val {
		val, count = left.val, left.count+right.count
	} else if left.count >= right.count {
		val, count = left.val, left.count-right.count
	} else {
		val, count = right.val, right.count-left.count
	}
	return &SegTree{val, count}
}

func querySegTree(segT []*SegTree, idx, start, end, left, right int, result *SegTree) *SegTree {
	if idx >= len(segT) {
		return result
	}
	if left <= start && end <= right {
		return combineSegTreeChild(result, segT[idx])
	}
	mid := start + (end-start)/2
	if left <= mid {
		result = querySegTree(segT, idx*2, start, mid, left, right, result)
	}
	if right >= mid {
		result = querySegTree(segT, idx*2+1, mid+1, end, left, right, result)
	}
	return result
}
