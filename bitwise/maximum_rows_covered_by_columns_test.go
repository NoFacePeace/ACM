package bitwise

import "testing"

func Test_maximumRows(t *testing.T) {
	type args struct {
		matrix    [][]int
		numSelect int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				matrix:    [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}},
				numSelect: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRows(tt.args.matrix, tt.args.numSelect); got != tt.want {
				t.Errorf("maximumRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
