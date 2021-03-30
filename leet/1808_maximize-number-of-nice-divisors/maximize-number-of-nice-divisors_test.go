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
		{5, 6},
		{8, 18},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := maxNiceDivisors(tt.n); got != tt.want {
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
