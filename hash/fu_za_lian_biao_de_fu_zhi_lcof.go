package hash

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	root := &Node{}
	m := map[*Node]*Node{}
	n1 := head
	n2 := root
	for n1 != nil {
		node := &Node{
			Val: n1.Val,
		}
		m[n1] = node
		n2.Next = node
		n2 = node
		n1 = n1.Next
	}
	n2 = root.Next
	n1 = head
	for n1 != nil {
		n2.Random = m[n1.Random]
		n2 = n2.Next
		n1 = n1.Next
	}
	return root.Next
}
