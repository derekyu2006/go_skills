package go_slice

import "testing"

func TestMainFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "slice_func_args",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MainFunc()
		})
	}
}
