package array

import "testing"

func Test_countDifferentSubsequenceGCDs(t *testing.T) {
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
				nums: []int{6, 10, 3},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDifferentSubsequenceGCDs(tt.args.nums); got != tt.want {
				t.Errorf("countDifferentSubsequenceGCDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
