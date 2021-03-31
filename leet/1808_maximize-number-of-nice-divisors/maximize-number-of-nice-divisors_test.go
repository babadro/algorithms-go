package _1808_maximize_number_of_nice_divisors

import (
	"fmt"
	"testing"
)

func Test_maxNiceDivisors(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 6},
		{6, 9},
		{7, 12},
		{8, 18},
		{9, 27},
		{18, 729},
		{64, 947137513},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := maxNiceDivisors2(tt.n); got != tt.want {
				t.Errorf("maxNiceDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxNiceDivisorsBruteForce(t *testing.T) {
	for i := 1; i < 20; i++ {
		t.Log(maxNiceDivisorsBruteForce(i))
	}
}
