package dp

import "testing"

func Test_maxSumTwoNoOverlap(t *testing.T) {
	type args struct {
		nums      []int
		firstLen  int
		secondLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		nums:      []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
		// 		firstLen:  4,
		// 		secondLen: 3,
		// 	},
		// },
		{
			args: args{
				nums:      []int{1, 16, 20, 13, 2, 11, 12, 18, 8, 5, 18, 15, 0, 7},
				firstLen:  3,
				secondLen: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumTwoNoOverlap(tt.args.nums, tt.args.firstLen, tt.args.secondLen); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
