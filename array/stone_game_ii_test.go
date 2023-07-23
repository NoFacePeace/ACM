package array

import "testing"

func Test_stoneGameII(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				piles: []int{1, 2, 3, 4, 5, 100},
			},
			want: 104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameII(tt.args.piles); got != tt.want {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}
