package _050_pow

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	cases := []struct {
		x float64
		n int
	}{
		{2, 0},
		{-2, 0},
		{1, 1243123423},
		{-2, 3},
		{-1, 1243123423},
		{-1, 1243123423},
		{2.1, 3},
		{2, -2},
		{1.00000, -2147483648},
	}

	for i, c := range cases {
		expected := math.Pow(c.x, float64(c.n))
		if fact := myPow(c.x, c.n); fact != expected {
			t.Errorf("case#%d, want %f, got %f", i+1, expected, fact)
		}
	}
}
