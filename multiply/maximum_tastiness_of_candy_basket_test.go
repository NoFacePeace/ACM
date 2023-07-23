package multiply

import "testing"

func Test_maximumTastiness(t *testing.T) {
	type args struct {
		price []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		price: []int{34, 116, 83, 15, 150, 56, 69, 42, 26},
		// 		k:     6,
		// 	},
		// 	want: 19,
		// },
		// {
		// 	args: args{
		// 		price: []int{13, 5, 1, 8, 21, 2},
		// 		k:     3,
		// 	},
		// 	want: 8,
		// },
		{
			args: args{
				price: []int{144, 69, 103, 148, 184, 50, 129, 154, 2},
				k:     4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTastiness(tt.args.price, tt.args.k); got != tt.want {
				t.Errorf("maximumTastiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
