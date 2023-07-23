package firstsearch

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	args: args{
		// 		graph: [][]int{{1, 2}, {3}, {3}, {}},
		// 	},
		// },
		// {
		// 	args: args{
		// 		graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
		// 	},
		// },
		{
			args: args{
				graph: [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
