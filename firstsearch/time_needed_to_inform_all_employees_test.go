package firstsearch

import "testing"

func Test_numOfMinutes(t *testing.T) {
	type args struct {
		n          int
		headID     int
		manager    []int
		informTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:          15,
				headID:     0,
				manager:    []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6},
				informTime: []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfMinutes(tt.args.n, tt.args.headID, tt.args.manager, tt.args.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
