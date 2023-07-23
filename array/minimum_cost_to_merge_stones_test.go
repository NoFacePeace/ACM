package array

import "testing"

func Test_mergeStones(t *testing.T) {
	type args struct {
		stones []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				stones: []int{3, 2, 4, 1},
				k:      2,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStones(tt.args.stones, tt.args.k); got != tt.want {
				t.Errorf("mergeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
