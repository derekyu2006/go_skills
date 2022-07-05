package go_syncmap

import "testing"

func TestSyncMapExecs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_syncmap",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SyncMapExecs()
		})
	}
}
