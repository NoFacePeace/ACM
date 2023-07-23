package array

import "testing"

func Test_maxHeight(t *testing.T) {
	type args struct {
		cuboids [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				cuboids: [][]int{
					{74, 7, 80},
					{7, 52, 61},
					{62, 41, 37},
					{91, 58, 26},
					{88, 98, 5},
					{72, 93, 23},
					{56, 58, 94},
					{88, 8, 64},
					{32, 55, 5},
				},
			},
			want: 301,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHeight(tt.args.cuboids); got != tt.want {
				t.Errorf("maxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
