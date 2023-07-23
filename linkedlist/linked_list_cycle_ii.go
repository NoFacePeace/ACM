package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}
	first := head
	second := head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			break
		}
	}
	if second == nil {
		return nil
	}
	if second.Next == nil {
		return nil
	}
	second = head
	for second != first {
		first = first.Next
		second = second.Next
	}
	return second
}
