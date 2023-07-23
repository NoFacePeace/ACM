package array

import "testing"

func Test_splitArraySameAverage(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{6, 8, 18, 3, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArraySameAverage(tt.args.nums); got != tt.want {
				t.Errorf("splitArraySameAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
