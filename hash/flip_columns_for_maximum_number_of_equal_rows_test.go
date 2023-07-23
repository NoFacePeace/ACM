package hash

import "testing"

func Test_maxEqualRowsAfterFlips(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				matrix: [][]int{{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1}, {1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}, {1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1}, {1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}, {1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1}},
			},
			want: 2,
		},
		// {
		// 	args: args{
		// 		matrix: [][]int{{0, 1}, {1, 0}},
		// 	},
		// 	want: 2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEqualRowsAfterFlips(tt.args.matrix); got != tt.want {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
