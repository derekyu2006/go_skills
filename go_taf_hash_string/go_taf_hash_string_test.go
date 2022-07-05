package go_taf_hash_string

import (
	"testing"

	"github.com/derekyu2006/dglog"
)

func Test_calcMagicHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "test_go_taf_hash_string",
			args: args{
				s: "09f09f53b160101af3b78da9095d88cb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcMagicHash(tt.args.s)
			dglog.Debugf("calcMagicHash: %d", got)
		})
	}
}
