package array

import "testing"

func Test_makeArrayIncreasing(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr1: []int{0, 11, 6, 1, 4, 3},
				arr2: []int{5, 4, 11, 10, 1, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeArrayIncreasing(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("makeArrayIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
