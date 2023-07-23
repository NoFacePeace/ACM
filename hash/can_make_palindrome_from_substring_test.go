package hash

import (
	"reflect"
	"testing"
)

func Test_canMakePaliQueries(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				s: "berilmre",
				queries: [][]int{
					{1, 6, 2},
					{6, 7, 9},
					{7, 7, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakePaliQueries(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePaliQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
