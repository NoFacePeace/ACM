package str

import "testing"

func Test_printBin(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num: 0.1,
			},
			want: "ERROR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printBin(tt.args.num); got != tt.want {
				t.Errorf("printBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
