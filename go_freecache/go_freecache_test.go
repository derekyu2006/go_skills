package go_freecache

import "testing"

func TestFreeCacheDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_freecache",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FreeCacheDemo()
		})
	}
}