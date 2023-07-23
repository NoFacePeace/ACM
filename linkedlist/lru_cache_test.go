package linkedlist

import "testing"

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		m        map[int]*item
		head     *item
		tail     *item
		size     int
		capacity int
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LRUCache{
				m:        tt.fields.m,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
			}
			this.Put(tt.args.key, tt.args.value)
		})
	}
}
