package bitwise

import "testing"

func Test_minNumberOfSemesters(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:         4,
				relations: [][]int{{2, 1}, {3, 1}, {1, 4}},
				k:         2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfSemesters(tt.args.n, tt.args.relations, tt.args.k); got != tt.want {
				t.Errorf("minNumberOfSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
