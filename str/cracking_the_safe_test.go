package str

import "testing"

func Test_crackSafe(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				n: 1,
				k: 10,
			},
			want: "01100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crackSafe(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("crackSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
