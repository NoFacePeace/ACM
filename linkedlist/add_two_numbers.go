package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var head *ListNode
	carry := 0
	next := head
	for l1 != nil && l2 != nil {
		n := l1.Val + l2.Val + carry
		tmp := &ListNode{
			Val: n % 10,
		}
		carry = n / 10
		if head == nil {
			head = tmp
			next = tmp
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		next.Next = tmp
		next = tmp
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		n := l1.Val + carry
		tmp := &ListNode{
			Val: n % 10,
		}
		carry = n / 10
		next.Next = tmp
		next = tmp
		l1 = l1.Next
	}
	for l2 != nil {
		n := l2.Val + carry
		tmp := &ListNode{
			Val: n % 10,
		}
		carry = n / 10
		next.Next = tmp
		next = tmp
		l2 = l2.Next
	}
	if carry != 0 {
		tmp := &ListNode{
			Val: carry,
		}
		next.Next = tmp
	}
	return head
}
