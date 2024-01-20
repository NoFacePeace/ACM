package set

import "github.com/emirpasic/gods/trees/redblacktree"

type StockPrice struct {
	tree *redblacktree.Tree
	m    map[int]int
	now  int
}

func Constructor() StockPrice {
	return StockPrice{
		tree: redblacktree.NewWithIntComparator(),
		m:    map[int]int{},
		now:  0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if _, ok := this.m[timestamp]; ok {
		p := this.m[timestamp]
		cnt, _ := this.tree.Get(p)
		if cnt == 1 {
			this.tree.Remove(p)
		} else {
			this.tree.Put(p, cnt.(int)-1)
		}
	}
	this.m[timestamp] = price
	if cnt, ok := this.tree.Get(price); ok {
		this.tree.Put(price, cnt.(int)+1)
	} else {
		this.tree.Put(price, 1)
	}
	if timestamp > this.now {
		this.now = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.m[this.now]
}

func (this *StockPrice) Maximum() int {
	node := this.tree.Right()
	key := node.Key.(int)
	return key
}

func (this *StockPrice) Minimum() int {
	node := this.tree.Left()
	key := node.Key.(int)
	return key
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
