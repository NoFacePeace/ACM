package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
		f   func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{5, 6, 7, 2, 1, 0},
				f: func(a, b int) bool {
					return a > b
				},
			},
			want: []int{0, 1, 2, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
		f   func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
				f: func(a, b int) bool {
					return a > b
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
