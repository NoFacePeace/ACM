package stack

type StockSpanner struct {
	cnt     int
	indexes []int
	prices  []int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{
		cnt:     0,
		indexes: []int{},
		prices:  []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.cnt++
	for len(this.prices) > 0 {
		n := len(this.prices)
		if this.prices[n-1] > price {
			break
		}
		this.prices = this.prices[:n-1]
		this.indexes = this.indexes[:n-1]
	}
	n := len(this.prices)
	span := this.cnt
	if n != 0 {
		span = this.cnt - this.indexes[n-1]
	}
	this.indexes = append(this.indexes, this.cnt)
	this.prices = append(this.prices, price)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
