package stack

type MinStack struct {
	dataStack []int
	minStack  []int
	min       int
}

func Constructor() MinStack {
	return MinStack{
		dataStack: []int{},
		minStack:  []int{},
		min:       0,
	}
}

func (this *MinStack) Push(val int) {
	this.dataStack = append(this.dataStack, val)
	if len(this.minStack) == 0 {
		this.min = val
		this.minStack = append(this.minStack, val)
		return
	}
	if val > this.min {
		this.minStack = append(this.minStack, this.min)
		return
	}
	this.min = val
	this.minStack = append(this.minStack, this.min)
}

func (this *MinStack) Pop() {
	if len(this.dataStack) == 0 {
		return
	}
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
	if len(this.minStack) == 0 {
		return
	}
	this.min = this.minStack[len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
