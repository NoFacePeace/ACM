package array

import "testing"

func Test_minScoreTriangulation(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		values: []int{4, 3, 4, 3, 5},
		// 	},
		// 	want: 135,
		// },
		// {
		// 	args: args{
		// 		values: []int{2, 2, 2, 2, 1},
		// 	},
		// 	want: 12,
		// },
		// {
		// 	args: args{
		// 		values: []int{1,3,1,4,1,5},
		// 	},
		// 	want: 13,
		// },
		{
			args: args{
				values: []int{5, 3, 5, 5, 1, 6, 2, 3},
			},
			want: 88,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScoreTriangulation(tt.args.values); got != tt.want {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
