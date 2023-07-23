package tree

import "testing"

func TestCountIntervals_Add(t *testing.T) {
	type fields struct {
		root *item
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			args: args{
				left:  7,
				right: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CountIntervals{
				root: tt.fields.root,
			}
			this.Add(2, 3)
			this.Add(tt.args.left, tt.args.right)
			this.Add(5, 8)
			this.Count()
		})
	}
}
