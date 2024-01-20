package set

import "github.com/emirpasic/gods/trees/redblacktree"

type CountIntervals struct {
	tree  *redblacktree.Tree
	count int
}

func NewCountIntervals() CountIntervals {
	return CountIntervals{
		tree:  redblacktree.NewWithIntComparator(),
		count: 0,
	}
}

func (this *CountIntervals) Add(left int, right int) {
	if node, ok := this.tree.Floor(left); ok && node.Value.(int) >= left {
		left = node.Key.(int)
		right = max(right, node.Value.(int))
		this.count -= node.Value.(int) - node.Key.(int) + 1
		this.tree.Remove(node.Key)
	}
	for node, ok := this.tree.Ceiling(left); ok && node.Key.(int) <= right; node, ok = this.tree.Ceiling(left) {
		this.count -= node.Value.(int) - node.Key.(int) + 1
		right = max(right, node.Value.(int))
		this.tree.Remove(node.Key)
	}
	this.count += right - left + 1
	this.tree.Put(left, right)
}

func (this *CountIntervals) Count() int {
	return this.count
}
