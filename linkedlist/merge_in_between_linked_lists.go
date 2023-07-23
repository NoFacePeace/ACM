package linkedlist

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = list1
	prev := head
	cnt := 0
	for cnt != a {
		prev = prev.Next
		cnt++
	}
	aList := prev.Next
	prev.Next = list2
	for prev.Next != nil {
		prev = prev.Next
	}
	bList := aList
	for cnt != b {
		bList = bList.Next
		cnt++
	}
	prev.Next = bList.Next
	return head.Next
}
