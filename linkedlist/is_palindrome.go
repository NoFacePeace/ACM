package linkedlist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l := reverseList(slow.Next)
	for l != nil {
		if head.Val == l.Val {
			head = head.Next
			l = l.Next
		} else {
			return false
		}
	}
	return true
}
