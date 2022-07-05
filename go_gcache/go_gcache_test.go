package go_gcache

import "testing"

func TestGCacheOne(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_gcache_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GCacheOne()
		})
	}
}
