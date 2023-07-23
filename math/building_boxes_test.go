package math

import "testing"

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 126,
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBoxes(tt.args.n); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
