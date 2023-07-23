package array

import (
	"reflect"
	"testing"
)

func Test_qSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
