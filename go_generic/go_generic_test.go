package go_generic

import (
	"reflect"
	"testing"

	"github.com/derekyu2006/dglog"
)

func TestGetMaxNum(t *testing.T) {
	type args[T int32 | int8] struct {
		a T
		b T
	}

	type TestWrapper[T int32 | int8] struct {
		name string
		args args[int32]
		want T
	}

	tests := []TestWrapper[int32]{
		{
			name: "test_go_generic",
			args: args[int32]{
				a: 3,
				b: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dglog.Debugf("a, b: %+v, %+v", tt.args.a, tt.args.b)
			if got := GetMaxNum(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaxNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
