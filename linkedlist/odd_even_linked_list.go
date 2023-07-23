package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h1 := head
	h2 := head.Next
	odd := h1
	even := h2
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		if odd.Next == nil {
			even.Next = nil
			break
		}
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = h2
	return head
}
