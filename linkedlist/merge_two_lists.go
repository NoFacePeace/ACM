package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	next := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if head == nil {
				head = list1
				next = list1
			} else {
				next.Next = list1
				next = list1
			}
			list1 = list1.Next
		} else {
			if head == nil {
				head = list2
				next = list2
			} else {
				next.Next = list2
				next = list2
			}
			list2 = list2.Next
		}
	}
	if list1 != nil {
		if head == nil {
			head = list1
		} else {
			next.Next = list1
		}
	}
	if list2 != nil {
		if head == nil {
			head = list2
		} else {
			next.Next = list2
		}
	}
	return head
}
