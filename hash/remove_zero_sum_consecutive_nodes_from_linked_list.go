package hash

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	ans := &ListNode{}
	ans.Next = head
	next := ans
	ps := 0
	m := map[int]*ListNode{}
	for next != nil {
		ps += next.Val
		m[ps] = next
		next = next.Next
	}
	next = ans
	ps = 0
	for next != nil {
		ps += next.Val
		next.Next = m[ps].Next
		next = next.Next
	}
	return ans.Next
}
