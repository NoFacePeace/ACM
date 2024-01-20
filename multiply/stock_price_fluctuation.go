package multiply

import "container/heap"

type StockPrice struct {
	now     int
	m       map[int]int
	minheap *MinHeapStock
	maxheap *MinHeapStock
}

type stock struct {
	timestamp int
	price     int
}

type MinHeapStock []stock

func (h MinHeapStock) Len() int {
	return len(h)
}

func (h MinHeapStock) Less(a, b int) bool {
	return h[a].price < h[b].price
}

func (h MinHeapStock) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *MinHeapStock) Push(x interface{}) {
	*h = append(*h, x.(stock))
}

func (h *MinHeapStock) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeapStock) Top() stock {
	return h[0]
}

func (h MinHeapStock) Tail() stock {
	return h[h.Len()-1]
}

func ConstructorStockPrice() StockPrice {
	return StockPrice{
		m:       map[int]int{},
		minheap: &MinHeapStock{},
		maxheap: &MinHeapStock{},
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.m[timestamp] = price
	if timestamp > this.now {
		this.now = timestamp
	}
	heap.Push(this.minheap, stock{
		timestamp: timestamp,
		price:     price,
	})
	heap.Push(this.maxheap, stock{
		timestamp: timestamp,
		price:     -price,
	})
}

func (this *StockPrice) Current() int {
	return this.m[this.now]
}

func (this *StockPrice) Maximum() int {
	s := this.maxheap.Top()
	for s.price != -this.m[s.timestamp] {
		heap.Pop(this.maxheap)
		s = this.maxheap.Top()
	}
	return -s.price
}

func (this *StockPrice) Minimum() int {
	s := this.minheap.Top()
	for s.price != this.m[s.timestamp] {
		heap.Pop(this.minheap)
		s = this.minheap.Top()
	}
	return s.price
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
