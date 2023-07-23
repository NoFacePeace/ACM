package array

import "testing"

func Test_isGoodArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{12, 5, 7, 23},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoodArray(tt.args.nums); got != tt.want {
				t.Errorf("isGoodArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
