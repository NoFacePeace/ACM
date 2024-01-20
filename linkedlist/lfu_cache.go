package linkedlist

type LFUCache struct {
	head     *LFUBucket
	bm       map[int]*LFUBucket
	nm       map[int]*LFUNode
	size     int
	capacity int
}

type LFUBucket struct {
	last *LFUBucket
	next *LFUBucket
	head *LFUNode
	cnt  int
	size int
}

type LFUNode struct {
	last  *LFUNode
	next  *LFUNode
	key   int
	value int
	cnt   int
}

func Constructor(capacity int) LFUCache {
	bucket := &LFUBucket{}
	bucket.last = bucket
	bucket.next = bucket
	cache := LFUCache{
		head:     bucket,
		bm:       map[int]*LFUBucket{},
		nm:       map[int]*LFUNode{},
		size:     0,
		capacity: capacity,
	}
	return cache
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.nm[key]
	if !ok {
		return -1
	}
	this.inc(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.nm[key]
	if ok {
		node.value = value
		this.inc(node)
		return
	}
	if this.size == this.capacity {
		this.del()
	}
	node = &LFUNode{
		key:   key,
		value: value,
		cnt:   1,
	}
	this.nm[key] = node
	this.size++
	bucket, ok := this.bm[node.cnt]
	if !ok {
		bucket = this.createBucket(node.cnt)
		this.insertBucket(this.head, bucket)
	}
	this.insertNode(node)
}

func (this *LFUCache) inc(node *LFUNode) {
	this.removeNode(node)
	ob := this.bm[node.cnt]
	node.cnt++
	nb, ok := this.bm[node.cnt]
	if !ok {
		nb = this.createBucket(node.cnt)
		this.insertBucket(ob, nb)
	}
	this.insertNode(node)
	if ob.size == 0 {
		this.delBucket(ob)
	}
}

func (this *LFUCache) del() {
	bucket := this.head.last
	last := bucket.head.last
	this.delNode(last)
	if bucket.size == 0 {
		this.delBucket(bucket)
	}
}

func (this *LFUCache) delNode(node *LFUNode) {
	node.next.last = node.last
	node.last.next = node.next
	delete(this.nm, node.key)
	this.bm[node.cnt].size--
	this.size--
}

func (this *LFUCache) removeNode(node *LFUNode) {
	node.next.last = node.last
	node.last.next = node.next
	bucket := this.bm[node.cnt]
	bucket.size--
}

func (this *LFUCache) createBucket(cnt int) *LFUBucket {
	bucket := &LFUBucket{}
	bucket.cnt = cnt
	bucket.head = &LFUNode{}
	bucket.head.next = bucket.head
	bucket.head.last = bucket.head
	this.bm[cnt] = bucket
	return bucket
}

func (this *LFUCache) insertBucket(old, new *LFUBucket) {
	last := old.last
	last.next = new
	new.last = last
	new.next = old
	old.last = new
}

func (this *LFUCache) insertNode(node *LFUNode) {
	bucket := this.bm[node.cnt]
	next := bucket.head.next
	bucket.head.next = node
	next.last = node
	node.last = bucket.head
	node.next = next
	bucket.size++
}

func (this *LFUCache) delBucket(bucket *LFUBucket) {
	bucket.last.next = bucket.next
	bucket.next.last = bucket.last
	delete(this.bm, bucket.cnt)
}
