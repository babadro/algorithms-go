package _33

import (
	"fmt"
	"math"
	"testing"
)

func Test_sum(t *testing.T) {
	t.Log(sum(1_000_000_000, 500_000_000, 1_000_000_000))
	t.Log(math.MaxInt64)
}

func Test_maxValue(t *testing.T) {
	tests := []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		{
			n: 4, index: 2, maxSum: 6, want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n: %d, index: %d, maxSum: %d", tt.n, tt.index, tt.maxSum), func(t *testing.T) {
			if got := maxValue(tt.n, tt.index, tt.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
