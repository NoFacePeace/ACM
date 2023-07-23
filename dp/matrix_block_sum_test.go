package dp

import (
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixBlockSum(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
