package backtrack

import "testing"

func Test_maxScoreWords(t *testing.T) {
	type args struct {
		words   []string
		letters []byte
		score   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words:   []string{"leetcode"},
				letters: []byte{'l', 'e', 't', 'c', 'o', 'd'},
				score:   []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreWords(tt.args.words, tt.args.letters, tt.args.score); got != tt.want {
				t.Errorf("maxScoreWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
