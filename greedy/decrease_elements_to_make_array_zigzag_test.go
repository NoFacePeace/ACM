package greedy

import "testing"

func Test_movesToMakeZigzag(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{7, 4, 8, 9, 7, 7, 5},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToMakeZigzag(tt.args.nums); got != tt.want {
				t.Errorf("movesToMakeZigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}
