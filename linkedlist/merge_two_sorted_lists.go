package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var ans *ListNode
	if list1.Val < list2.Val {
		ans = list1
		list1 = list1.Next
	} else {
		ans = list2
		list2 = list2.Next
	}
	next := ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next = list2
			list2 = list2.Next
		}
		next = next.Next
	}
	for list1 != nil {
		next.Next = list1
		list1 = list1.Next
		next = next.Next
	}
	for list2 != nil {
		next.Next = list2
		list2 = list2.Next
		next = next.Next
	}
	return ans
}
