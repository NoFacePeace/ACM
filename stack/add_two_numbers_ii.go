package stack

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := []int{}
	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}
	arr2 := []int{}
	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}
	n := len(arr1) - 1
	m := len(arr2) - 1
	carry := 0
	var next *ListNode
	for n >= 0 && m >= 0 {
		val := arr1[n] + arr2[m] + carry
		carry = val / 10
		val %= 10
		node := &ListNode{
			Val: val,
		}
		node.Next = next
		next = node
		n--
		m--
	}
	for n >= 0 {
		val := arr1[n] + carry
		carry = val / 10
		val %= 10
		node := &ListNode{
			Val: val,
		}
		node.Next = next
		next = node
		n--
	}
	for m >= 0 {
		val := arr2[m] + carry
		carry = val / 10
		val %= 10
		node := &ListNode{
			Val: val,
		}
		node.Next = next
		next = node
		m--
	}
	if carry != 0 {
		node := &ListNode{
			Val: carry,
		}
		node.Next = next
		next = node
	}
	return next
}
