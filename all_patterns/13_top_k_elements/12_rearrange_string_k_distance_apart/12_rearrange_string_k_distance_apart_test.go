package _12_rearrange_string_k_distance_apart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reorganizeString(t *testing.T) {
	tests := []struct {
		str  string
		k    int
		want bool
	}{
		{"mmpp", 2, true},
		{"Programming", 3, true},
		{"aab", 2, true},
		{"aappa", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := reorganizeString(tt.str, tt.k)
			if tt.want == false {
				require.Empty(t, got)
			} else {
				require.Len(t, got, len(tt.str))
			}
		})
	}
}
