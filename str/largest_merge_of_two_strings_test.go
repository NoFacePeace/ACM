package str

import "testing"

func Test_largestMerge(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				word1: "cabaa",
				word2: "bcaaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestMerge(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("largestMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
