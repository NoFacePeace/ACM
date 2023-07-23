package other

import (
	"reflect"
	"testing"
)

func Test_circularPermutation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:     3,
				start: 2,
			},
			want: []int{2, 6, 7, 5, 4, 0, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularPermutation(tt.args.n, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
