package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				l1: buildLinkedList([]int{2, 4, 3}),
				l2: buildLinkedList([]int{5, 6, 4}),
			},
			want: buildLinkedList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildLinkedList(arr []int) *ListNode {
	if arr == nil {
		return nil
	}
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{
		Val: arr[0],
	}
	next := head
	for i := 1; i < len(arr); i++ {
		tmp := &ListNode{
			Val: arr[i],
		}
		next.Next = tmp
		next = tmp
	}
	return head
}
