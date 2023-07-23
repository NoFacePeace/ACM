package array

import "testing"

func Test_storeWater(t *testing.T) {
	type args struct {
		bucket []int
		vat    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		bucket: []int{16, 29, 42, 70, 42, 9},
		// 		vat:    []int{89, 44, 50, 90, 94, 91},
		// 	},
		// 	want: 11,
		// },
		{
			args: args{
				bucket: []int{1, 3},
				vat:    []int{6, 8},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storeWater(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("storeWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
