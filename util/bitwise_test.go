package util

import "testing"

func Test_numberOfLeadingZeros(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 2,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLeadingZeros(tt.args.num); got != tt.want {
				t.Errorf("numberOfLeadingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
