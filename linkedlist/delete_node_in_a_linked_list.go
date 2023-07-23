package linkedlist

func deleteNode(node *ListNode) {
	var d func(node *ListNode)
	d = func(node *ListNode) {
		if node.Next.Next == nil {
			node.Val = node.Next.Val
			node.Next = nil
			return
		}
		node.Val = node.Next.Val
		d(node.Next)
	}
	d(node)
	return
}
