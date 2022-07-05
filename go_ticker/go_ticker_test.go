package go_ticker

import "testing"

func TestTickExec(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_ticker",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TickExec()
		})
	}
}
