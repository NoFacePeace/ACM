package stack

import (
	"reflect"
	"testing"
)

func Test_braceExpansionII(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				expression: "{{a,z},a{b,c},{ab,z}}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := braceExpansionII(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
