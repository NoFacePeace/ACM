package array

import "testing"

func Test_reachNumber(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				target: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachNumber(tt.args.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
