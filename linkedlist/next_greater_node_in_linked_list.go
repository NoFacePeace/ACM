package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	val := []int{}
	idx := []int{}
	i := 0
	m := map[int]int{}
	for head != nil {
		l := len(val)
		if l == 0 || val[l-1] >= head.Val {
			val = append(val, head.Val)
			idx = append(idx, i)
			i++
			head = head.Next
			continue
		}
		for l > 0 && val[l-1] < head.Val {
			m[idx[l-1]] = head.Val
			idx = idx[:l-1]
			val = val[:l-1]
			l = len(val)
		}
		val = append(val, head.Val)
		idx = append(idx, i)
		i++
		head = head.Next
	}
	for _, v := range idx {
		m[v] = 0
	}
	ans := make([]int, len(m))
	for k, v := range m {
		ans[k] = v
	}
	return ans
}
