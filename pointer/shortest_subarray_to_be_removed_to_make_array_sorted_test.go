package pointer

import (
	"testing"
)

func Test_findLengthOfShortestSubarray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		arr: []int{1, 2, 3, 10, 0, 7, 8, 9},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	args: args{
		// 		arr: []int{13, 0, 14, 7, 18, 18, 18, 16, 8, 15, 20},
		// 	},
		// 	want: 8,
		// },
		{
			args: args{
				arr: []int{1, 2, 3, 10, 4, 2, 3, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfShortestSubarray(tt.args.arr); got != tt.want {
				t.Errorf("findLengthOfShortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
