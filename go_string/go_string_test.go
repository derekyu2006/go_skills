package go_string

import "testing"

func TestMainFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_func_string_args",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MainFunc()
		})
	}
}
