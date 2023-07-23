package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}
	var h *Node
	var next *Node
	for head != nil {
		if h == nil {
			h = &Node{
				Val: head.Val,
			}
			next = h
			m[head] = h
			if head.Random == nil {
				head = head.Next
				continue
			}
			if head.Random == head {
				h.Random = h
				head = head.Next
				continue
			}
			r := &Node{
				Val: head.Random.Val,
			}
			m[head.Random] = r
			h.Random = r
			head = head.Next
			continue
		}
		var node *Node
		if _, ok := m[head]; ok {
			node = m[head]
		} else {
			node = &Node{
				Val: head.Val,
			}
			m[head] = node
		}

		next.Next = node
		next = next.Next
		if head.Random == nil {
			head = head.Next
			continue
		}
		var r *Node
		if _, ok := m[head.Random]; ok {
			r = m[head.Random]
		} else {
			r = &Node{
				Val: head.Random.Val,
			}
			m[head.Random] = r
		}
		node.Random = r
		head = head.Next
	}
	return h
}
