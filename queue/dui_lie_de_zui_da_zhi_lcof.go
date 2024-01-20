package queue

type MaxQueue struct {
	q  []int
	pq []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:  []int{},
		pq: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	n := len(this.q)
	if n == 0 {
		return -1
	}
	return this.pq[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	i := 0
	for i < len(this.pq) {
		if this.pq[i] < value {
			break
		}
		i++
	}
	this.pq = this.pq[:i]
	this.pq = append(this.pq, value)
}

func (this *MaxQueue) Pop_front() int {
	n := len(this.q)
	if n == 0 {
		return -1
	}
	x := this.q[0]
	if x == this.pq[0] {
		this.pq = this.pq[1:]
	}
	this.q = this.q[1:]
	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
