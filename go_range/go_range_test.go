package go_range

import "testing"

func TestRangeNil(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_rangenil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeNil()
		})
	}
}
