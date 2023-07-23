package array

import (
	"reflect"
	"testing"
)

func Test_matrixRankTransform(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				matrix: [][]int{
					{-37, -50, -3, 44},
					{-37, 46, 13, -32},
					{47, -42, -3, -40},
					{-17, -22, -39, 24},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixRankTransform(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
