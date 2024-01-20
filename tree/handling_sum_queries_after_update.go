package tree

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	t := NewSegTree(nums1)
	sum := 0
	n := len(nums1)
	for _, v := range nums2 {
		sum += v
	}
	ans := []int64{}
	for _, q := range queries {
		if q[0] == 1 {
			t.modify(0, q[1], q[2])
			continue
		}
		if q[0] == 2 {
			sum += t.query(0, 0, n-1) * q[1]
			continue
		}
		ans = append(ans, int64(sum))
	}
	return ans
}

type SegNode struct {
	l, r, sum int
	lazy      bool
}

type SegTree []SegNode

func NewSegTree(nums []int) *SegTree {
	t := SegTree{}
	n := len(nums)
	t = append(t, make([]SegNode, 4*n)...)
	t.build(0, 0, n-1, nums)
	return &t
}

func (t SegTree) build(id, l, r int, nums []int) {
	t[id].l = l
	t[id].r = r
	if l == r {
		t[id].sum = nums[l]
		return
	}
	mid := (l + r) / 2
	t.build(id*2+1, l, mid, nums)
	t.build(id*2+2, mid+1, r, nums)
	t[id].sum = t[id*2+1].sum + t[id*2+2].sum
}

func (t SegTree) query(id, l, r int) int {
	if t[id].l >= l && t[id].r <= r {
		return t[id].sum
	}
	if t[id].l > r {
		return 0
	}
	if t[id].r < l {
		return 0
	}
	t.updateChild(id)
	sum := 0
	if t[id*2+1].r >= l {
		sum += t.query(id*2+1, l, r)
	}
	if t[id*2+2].l <= r {
		sum += t.query(id*2+2, l, r)
	}
	return sum
}

func (t SegTree) modify(id, l, r int) {
	if t[id].l >= l && t[id].r <= r {
		t.update(id)
		return
	}
	t.updateChild(id)
	if t[id*2+1].r >= l {
		t.modify(id*2+1, l, r)
	}
	if t[id*2+2].l <= r {
		t.modify(id*2+2, l, r)
	}
	t[id].sum = t[id*2+1].sum + t[id*2+2].sum
}

func (t SegTree) updateChild(id int) {
	if !t[id].lazy {
		return
	}
	t[id].lazy = false
	t.update(id*2 + 1)
	t.update(id*2 + 2)
}

func (t SegTree) update(id int) {
	t[id].sum = t[id].r - t[id].l + 1 - t[id].sum
	t[id].lazy = !t[id].lazy
}
