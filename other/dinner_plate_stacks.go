package other

type DinnerPlates struct {
	idle     []int
	arr      [][]int
	capacity int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
	}
}

func (this *DinnerPlates) Push(val int) {
	if len(this.idle) == 0 {
		this.arr = append(this.arr, []int{})
		n := len(this.arr)
		this.arr[n-1] = append(this.arr[n-1], val)
		if len(this.arr[n-1]) < this.capacity {
			this.idle = append(this.idle, n-1)
		}
		return
	}
	idx := this.idle[0]
	this.arr[idx] = append(this.arr[idx], val)
	if len(this.arr[idx]) < this.capacity {
		return
	}
	this.idle = this.idle[1:]
}

func (this *DinnerPlates) Pop() int {
	n := len(this.arr)
	for i := n - 1; i >= 0; i-- {
		m := len(this.arr[i])
		if m == 0 {
			this.arr = this.arr[:i]
			this.idle = this.idle[:len(this.idle)-1]
			continue
		}
		val := this.arr[i][m-1]
		this.arr[i] = this.arr[i][:m-1]
		if m == this.capacity {
			this.idle = append(this.idle, i)
		}
		return val
	}
	return -1
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if len(this.arr) <= index {
		return -1
	}
	m := len(this.arr[index])
	if m == 0 {
		return -1
	}
	val := this.arr[index][m-1]
	this.arr[index] = this.arr[index][:m-1]
	if m != this.capacity {
		return val
	}
	l := len(this.idle)
	this.idle = append(this.idle, index)
	for i := l - 1; i >= 0; i-- {
		if this.idle[i] < this.idle[i+1] {
			break
		}
		this.idle[i], this.idle[i+1] = this.idle[i+1], this.idle[i]
	}
	return val
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
