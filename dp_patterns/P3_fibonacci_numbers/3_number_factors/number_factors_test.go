package __number_factors

import (
	"fmt"
	"testing"
)

func TestCountWays(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 4},
		{5, 6},
		{15, 714},
		{25, 87841},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			if got := CountWaysOptimized(tt.n); got != tt.want {
				t.Errorf("CountWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
