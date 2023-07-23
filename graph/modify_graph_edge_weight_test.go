package graph

import (
	"reflect"
	"testing"
)

func Test_modifiedGraphEdges(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
		target      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	args: args{
		// 		n:           5,
		// 		edges:       [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}},
		// 		source:      0,
		// 		destination: 1,
		// 		target:      5,
		// 	},
		// },
		// {
		// 	args: args{
		// 		n:           3,
		// 		edges:       [][]int{{0, 1, -1}, {0, 2, 5}},
		// 		source:      0,
		// 		destination: 2,
		// 		target:      6,
		// 	},
		// },
		{
			args: args{
				n:           4,
				edges:       [][]int{{0, 1, -1}, {1, 2, -1}, {3, 1, -1}, {3, 0, 2}, {0, 2, 5}},
				source:      2,
				destination: 3,
				target:      8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedGraphEdges(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedGraphEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
