package tree

import (
	"reflect"
	"testing"
)

func Test_handleQuery(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{
				nums1:   []int{0, 1, 0, 0, 0, 0},
				nums2:   []int{14, 4, 13, 13, 47, 18},
				queries: [][]int{{3, 0, 0}, {1, 4, 4}, {1, 1, 4}, {1, 3, 4}, {3, 0, 0}, {2, 5, 0}, {1, 1, 3}, {2, 16, 0}, {2, 10, 0}, {3, 0, 0}, {3, 0, 0}, {2, 6, 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleQuery(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
