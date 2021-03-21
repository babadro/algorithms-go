package _33

import (
	"fmt"
	"testing"
)

func Test_maxValue(t *testing.T) {
	tests := []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		{n: 4, index: 2, maxSum: 6, want: 2},
		{6, 1, 10, 3},
		{1, 0, 3, 3},
		{3, 0, 6, 3},
		{3, 2, 6, 3},
		{3, 1, 7, 3},
		{3, 0, 3, 1},
		{3, 2, 3, 1},
		{3, 1, 3, 1},
		{3, 1, 4, 2},
		{5, 0, 28, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n: %d, index: %d, maxSum: %d", tt.n, tt.index, tt.maxSum), func(t *testing.T) {
			if got := maxValue(tt.n, tt.index, tt.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum1(t *testing.T) {
	tests := []struct {
		num   int
		n     int
		index int
		want  int
	}{
		{num: 2, n: 4, index: 2, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("num: %d, n: %d, index: %d", tt.num, tt.n, tt.index), func(t *testing.T) {
			if got := sum(tt.n, tt.index, tt.num); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
