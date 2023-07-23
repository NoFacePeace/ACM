package str

import "testing"

func Test_stringSort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "tea",
			},
			want: "aet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSort(tt.args.s); got != tt.want {
				t.Errorf("stringSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
