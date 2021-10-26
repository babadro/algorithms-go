package __staircase

import (
	"fmt"
	"testing"
)

func TestStairsBruteForce(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 4},
		{4, 7},
		{15, 5768},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			if got := StairsBottomUpOptimized(tt.n); got != tt.want {
				t.Errorf("StairsBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
