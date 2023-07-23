package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	head.Next = nil
	last := head
	for next != nil {
		tmp := next.Next
		next.Next = last
		last = next
		next = tmp
	}
	return last
}
