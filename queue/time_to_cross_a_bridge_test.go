package queue

import "testing"

func Test_findCrossingTime(t *testing.T) {
	type args struct {
		n    int
		k    int
		time [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		n:    1,
		// 		k:    3,
		// 		time: [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}},
		// 	},
		// },
		{
			args: args{
				n:    3,
				k:    2,
				time: [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCrossingTime(tt.args.n, tt.args.k, tt.args.time); got != tt.want {
				t.Errorf("findCrossingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
