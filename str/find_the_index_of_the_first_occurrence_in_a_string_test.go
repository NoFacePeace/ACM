package str

import (
	"reflect"
	"testing"
)

func Test_kmp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				s: "abcdabd",
			},
			want: []int{0, 0, 0, 0, 1, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
