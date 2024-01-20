package pointer

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
	}
	return false
}
