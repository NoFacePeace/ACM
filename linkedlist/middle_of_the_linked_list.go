package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	if second != nil {
		return first.Next
	}
	return first
}
