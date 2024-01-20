package linkedlist

import "testing"

func TestLFUCache_Get(t *testing.T) {
	type fields struct {
		head     *LFUBucket
		bm       map[int]*LFUBucket
		nm       map[int]*LFUNode
		size     int
		capacity int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			args: args{
				key: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(3)
			this.Put(2, 2)
			this.Put(1, 1)
			this.Get(2)
			this.Get(1)
			this.Get(2)
			this.Put(3, 3)
			this.Put(4, 4)
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("LFUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
