package algorithm

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		edges    [][]int
		maxMoves int
		n        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				edges:    [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}},
				maxMoves: 6,
				n:        3,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachableNodes(tt.args.edges, tt.args.maxMoves, tt.args.n); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
