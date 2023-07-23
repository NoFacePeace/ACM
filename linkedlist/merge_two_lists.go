package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var ans *ListNode
	var next *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if ans == nil {
				ans = list1
				next = list1
				list1 = list1.Next
				continue
			}
			next.Next = list1
			next = next.Next
			list1 = list1.Next
			continue
		}
		if ans == nil {
			ans = list2
			next = list2
			list2 = list2.Next
			continue
		}
		next.Next = list2
		next = next.Next
		list2 = list2.Next
	}
	if list1 != nil {
		next.Next = list1
	}
	if list2 != nil {
		next.Next = list2
	}
	return ans
}
