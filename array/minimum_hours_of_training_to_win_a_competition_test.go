package array

import "testing"

func Test_minNumberOfHours(t *testing.T) {
	type args struct {
		initialEnergy     int
		initialExperience int
		energy            []int
		experience        []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				initialEnergy:     5,
				initialExperience: 3,
				energy:            []int{1, 4, 3, 2},
				experience:        []int{2, 6, 3, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfHours(tt.args.initialEnergy, tt.args.initialExperience, tt.args.energy, tt.args.experience); got != tt.want {
				t.Errorf("minNumberOfHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
