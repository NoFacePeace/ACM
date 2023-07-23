package doublepointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{}
	h.Next = head
	next := h
	for i := 0; i < n; i++ {
		if next == nil {
			return h.Next
		}
		next = next.Next
	}
	del := h
	for next.Next != nil {
		del = del.Next
		next = next.Next
	}

	del.Next = del.Next.Next
	return h.Next
}
