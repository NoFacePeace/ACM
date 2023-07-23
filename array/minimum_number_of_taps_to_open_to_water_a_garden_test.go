package array

import "testing"

func Test_minTaps(t *testing.T) {
	type args struct {
		n      int
		ranges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:      3,
				ranges: []int{0, 0, 0, 0},
			},
			want: -11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTaps(tt.args.n, tt.args.ranges); got != tt.want {
				t.Errorf("minTaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
