package set

import "github.com/emirpasic/gods/sets/treeset"

type SmallestInfiniteSet struct {
	max  int
	tree *treeset.Set
}

func NewSmallestInfiniteSet() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		max:  1,
		tree: treeset.NewWithIntComparator(),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	num := this.max
	if this.tree.Empty() {
		this.max++
		return num
	}
	it := this.tree.Iterator()
	it.Next()
	num = it.Value().(int)
	this.tree.Remove(num)
	return num
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.max {
		this.tree.Add(num)
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
