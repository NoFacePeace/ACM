package sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr:   []int{3, 2, 1, 4},
				left:  0,
				right: 3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test 1",
			args: args{
				arr:   []int{3, 2, 4},
				left:  0,
				right: 2,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "test 2",
			args: args{
				arr:   []int{0, 3, -3, 4, -1},
				left:  0,
				right: 4,
			},
			want: []int{-3, -1, 0, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr:    []int{1, 2, 3},
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
