package math

import "testing"

func Test_minimumPerimeter(t *testing.T) {
	type args struct {
		neededApples int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				neededApples: 150,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPerimeter(tt.args.neededApples); got != tt.want {
				t.Errorf("minimumPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
