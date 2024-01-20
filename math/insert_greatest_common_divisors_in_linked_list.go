package math

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	next := head
	var gcd func(int, int) int
	gcd = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}
	for next.Next != nil {
		num := gcd(next.Val, next.Next.Val)
		node := &ListNode{
			Val: num,
		}
		last := next
		next = last.Next
		last.Next = node
		node.Next = next
	}
	return head
}
