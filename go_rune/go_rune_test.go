package go_rune

import "testing"

func TestRuneSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_rune_slice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuneSlice()
		})
	}
}
