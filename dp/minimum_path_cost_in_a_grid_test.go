package dp

import "testing"

func Test_minPathCost(t *testing.T) {
	type args struct {
		grid     [][]int
		moveCost [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid:     [][]int{{5, 3}, {4, 0}, {2, 1}},
				moveCost: [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathCost(tt.args.grid, tt.args.moveCost); got != tt.want {
				t.Errorf("minPathCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
