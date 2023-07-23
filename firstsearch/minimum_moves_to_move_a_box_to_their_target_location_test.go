package firstsearch

import "testing"

func Test_minPushBox(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]byte{{'#', '#', '#', '#', '#', '#'}, {'#', 'T', '.', '.', '#', '#'}, {'#', '.', '#', 'B', '.', '#'}, {'#', '.', '.', '.', '.', '#'}, {'#', '.', '.', '.', 'S', '#'}, {'#', '#', '#', '#', '#', '#'}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPushBox(tt.args.grid); got != tt.want {
				t.Errorf("minPushBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
