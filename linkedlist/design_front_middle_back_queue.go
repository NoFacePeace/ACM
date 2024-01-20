package linkedlist

type FrontMiddleBackQueue struct {
	first *FrontMiddleBackQueueNode
	last  *FrontMiddleBackQueueNode
	mid   *FrontMiddleBackQueueNode
	size  int
}

type FrontMiddleBackQueueNode struct {
	val  int
	last *FrontMiddleBackQueueNode
	next *FrontMiddleBackQueueNode
}

func NewFrontMiddleBackQueue() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		first: &FrontMiddleBackQueueNode{},
		last:  &FrontMiddleBackQueueNode{},
		mid:   &FrontMiddleBackQueueNode{},
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	node := &FrontMiddleBackQueueNode{
		val: val,
	}
	next := this.first.next
	node.next = next
	if next != nil {
		next.last = node
	}
	this.first.next = node
	node.last = this.first
	this.size++
	if this.size == 1 {
		this.last.next = node
		this.mid.next = node
		return
	}
	if this.size%2 == 0 {
		this.mid.next = this.mid.next.last
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.size == 0 {
		this.PushFront(val)
		return
	}
	node := &FrontMiddleBackQueueNode{
		val: val,
	}
	old := this.mid.next
	if this.size%2 == 0 {
		next := old.next
		old.next = node
		node.last = old
		node.next = next
		if next != nil {
			next.last = node
		}
	} else {
		last := old.last
		if last != nil {
			last.next = node
		}
		node.last = last
		node.next = old
		old.last = node
	}
	this.mid.next = node
	this.size++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	if this.size == 0 {
		this.PushFront(val)
		return
	}
	node := &FrontMiddleBackQueueNode{
		val: val,
	}
	old := this.last.next
	old.next = node
	node.last = old
	this.last.next = node
	if this.size%2 == 0 {
		this.mid.next = this.mid.next.next
	}
	this.size++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.size == 0 {
		return -1
	}
	if this.size == 1 {
		return this.PopLast()
	}
	if this.size%2 == 0 {
		this.mid.next = this.mid.next.next
	}
	node := this.first.next
	next := node.next
	this.first.next = next
	next.last = this.first
	this.size--
	return node.val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.size == 0 {
		return -1
	}
	if this.size == 1 {
		return this.PopLast()
	}
	node := this.mid.next
	last := node.last
	next := node.next
	last.next = next
	next.last = last
	if this.size%2 == 0 {
		this.mid.next = node.next
	} else {
		this.mid.next = node.last
	}
	this.size--
	return node.val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.size == 0 {
		return -1
	}
	if this.size == 1 {
		return this.PopLast()
	}
	if this.size%2 != 0 {
		this.mid.next = this.mid.next.last
	}
	node := this.last.next
	last := node.last
	last.next = nil
	this.last.next = last
	this.size--
	return node.val
}

func (this *FrontMiddleBackQueue) PopLast() int {
	node := this.first.next
	this.first.next = nil
	this.last.next = nil
	this.mid.next = nil
	this.size = 0
	return node.val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
