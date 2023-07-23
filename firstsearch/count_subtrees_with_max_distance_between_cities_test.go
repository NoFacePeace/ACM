package firstsearch

import (
	"reflect"
	"testing"
)

func Test_countSubgraphsForEachDiameter(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	args: args{
		// 		n:     4,
		// 		edges: [][]int{{1, 2}, {2, 3}, {2, 4}},
		// 	},
		// 	want: []int{3, 4, 0},
		// },
		{
			args: args{
				n:     4,
				edges: [][]int{{1, 3}, {1, 4}, {2, 3}},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubgraphsForEachDiameter(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubgraphsForEachDiameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
