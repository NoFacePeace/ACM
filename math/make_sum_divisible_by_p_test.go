package math

import "testing"

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2},
				p:    148,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
