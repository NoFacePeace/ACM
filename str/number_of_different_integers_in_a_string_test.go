package str

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				word: "a123bc34d8ef34",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDifferentIntegers(tt.args.word); got != tt.want {
				t.Errorf("numDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
