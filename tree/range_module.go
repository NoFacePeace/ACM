package tree

const N int = 1e9

type RangeModule struct {
	root *node
}
type node struct {
	left  *node
	right *node
	added bool
	lazy  bool
}

func NewRangeModule() RangeModule {
	return RangeModule{
		root: &node{},
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.update(this.root, 1, N, left, right-1, true)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.query(this.root, 1, N, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.update(this.root, 1, N, left, right-1, false)
}

func (this *RangeModule) update(n *node, l, r, i, j int, added bool) {
	if l >= i && r <= j {
		n.added = added
		n.lazy = true
		return
	}
	this.pushdown(n)
	m := (l + r) / 2
	if i <= m {
		this.update(n.left, l, m, i, j, added)
	}
	if j > m {
		this.update(n.right, m+1, r, i, j, added)
	}
	this.pushup(n)
}

func (this *RangeModule) query(n *node, l, r, i, j int) bool {
	if i <= l && j >= r {
		return n.added
	}
	this.pushdown(n)
	added := true
	m := (l + r) / 2
	if i <= m {
		added = added && this.query(n.left, l, m, i, j)
	}
	if j > m {
		added = added && this.query(n.right, m+1, r, i, j)
	}
	return added
}

func (this *RangeModule) pushup(n *node) {
	n.added = n.left.added && n.right.added
}

func (this *RangeModule) pushdown(n *node) {
	if n.left == nil {
		n.left = &node{}
	}
	if n.right == nil {
		n.right = &node{}
	}
	if n.lazy {
		n.left.added = n.added
		n.right.added = n.added
		n.left.lazy = true
		n.right.lazy = true
		n.lazy = false
	}
}
