package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{}
	h.Next = head
	next := h
	for next.Next != nil && next.Next.Next != nil {
		first := next.Next
		secord := next.Next.Next
		first.Next = secord.Next
		secord.Next = first
		next.Next = secord
		next = first
	}
	return h.Next
}
