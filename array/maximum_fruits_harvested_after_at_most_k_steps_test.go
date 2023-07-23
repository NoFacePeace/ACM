package array

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	type args struct {
		fruits   [][]int
		startPos int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				fruits:   [][]int{{2, 8}, {6, 3}, {8, 6}},
				startPos: 5,
				k:        4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalFruits(tt.args.fruits, tt.args.startPos, tt.args.k); got != tt.want {
				t.Errorf("maxTotalFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
