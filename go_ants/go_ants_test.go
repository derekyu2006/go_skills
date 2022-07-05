package go_ants

import "testing"

func TestAntsPool(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_antspool",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AntsPool()
		})
	}
}
