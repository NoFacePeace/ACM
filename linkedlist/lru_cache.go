package linkedlist

type LRUCache struct {
	m        map[int]*item
	head     *item
	tail     *item
	size     int
	capacity int
}

type item struct {
	Val  int
	Key  int
	Prev *item
	Next *item
}

func Constructor(capacity int) LRUCache {
	h := &item{}
	t := &item{}
	h.Next = t
	t.Prev = h
	return LRUCache{
		m:        map[int]*item{},
		head:     h,
		tail:     t,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}
	this.lru(key)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.Val = value
		this.lru(key)
		return
	}
	node = &item{
		Val: value,
		Key: key,
	}
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
	this.m[key] = node
	this.size++
	if this.size > this.capacity {
		delete(this.m, this.tail.Prev.Key)
		this.tail.Prev.Prev.Next = this.tail
		this.tail.Prev = this.tail.Prev.Prev
	}
}

func (this *LRUCache) lru(key int) {
	node := this.m[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}
