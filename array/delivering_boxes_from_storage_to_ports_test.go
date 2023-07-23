package array

import "testing"

func Test_boxDelivering(t *testing.T) {
	type args struct {
		boxes      [][]int
		portsCount int
		maxBoxes   int
		maxWeight  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				boxes: [][]int{
					{1, 2},
					{3, 3},
					{3, 1},
					{3, 1},
					{2, 4},
				},
				portsCount: 3,
				maxBoxes:   3,
				maxWeight:  6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boxDelivering(tt.args.boxes, tt.args.portsCount, tt.args.maxBoxes, tt.args.maxWeight); got != tt.want {
				t.Errorf("boxDelivering() = %v, want %v", got, tt.want)
			}
		})
	}
}
