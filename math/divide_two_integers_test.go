package math

import "testing"

func Test_quickAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				x: 3,
				y: 9,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickAdd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("quickAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
