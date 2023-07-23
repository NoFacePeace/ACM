package str

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "(name)is(age)yearsold",
				knowledge: [][]string{
					{"name", "bob"},
					{"age", "two"},
				},
			},
			want: "bobistwoyearsold",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
