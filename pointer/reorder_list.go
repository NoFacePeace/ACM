package pointer

func reorderList(head *ListNode) {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = head
	var dfs func(next *ListNode)
	dfs = func(next *ListNode) {
		if next == nil {
			return
		}
		dfs(next.Next)
		tmp := slow.Next
		slow.Next = next
		slow = slow.Next
		slow.Next = tmp
		slow = slow.Next
	}
	dfs(fast)
}
