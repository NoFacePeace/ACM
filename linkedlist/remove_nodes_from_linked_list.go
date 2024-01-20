package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	stack := []*ListNode{}
	for head != nil {
		for len(stack) > 0 {
			l := len(stack)
			if stack[l-1].Val >= head.Val {
				break
			}
			stack = stack[:l-1]
		}
		stack = append(stack, head)
		head = head.Next
	}
	var next *ListNode
	for k, v := range stack {
		if k == 0 {
			head = v
			v.Next = nil
			next = head
			continue
		}
		next.Next = v
		v.Next = nil
		next = v
	}
	return head
}
