package dp

import "testing"

func Test_maxTaxiEarnings(t *testing.T) {
	type args struct {
		n     int
		rides [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n:     20,
				rides: [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTaxiEarnings(tt.args.n, tt.args.rides); got != tt.want {
				t.Errorf("maxTaxiEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
