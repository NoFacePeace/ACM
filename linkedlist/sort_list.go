package linkedlist

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := []*ListNode{}
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	arr = sort(arr)
	var h *ListNode
	var next *ListNode
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			h = arr[i]
			next = h
			continue
		}
		next.Next = arr[i]
		next = next.Next
	}
	next.Next = nil
	return h
}

func sort(arr []*ListNode) []*ListNode {
	return qSort(arr, 0, len(arr)-1)
}

func qSort(arr []*ListNode, left, right int) []*ListNode {
	if left >= right {
		return arr
	}
	arr, p := partition(arr, left, right)
	arr = qSort(arr, left, p-1)
	arr = qSort(arr, p+1, right)
	return arr
}

func partition(arr []*ListNode, left, right int) ([]*ListNode, int) {
	mid := (left + right) / 2
	arr[right], arr[mid] = arr[mid], arr[right]
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j].Val < pivot.Val {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}
