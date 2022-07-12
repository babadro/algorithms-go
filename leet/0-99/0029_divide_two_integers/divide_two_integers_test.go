package _029_divide_two_integers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		dividend int
		divisor  int
		expected int
	}{
		{0, 1, 0},
		{1, 1, 1},
		{2, 1, 2},
		{5, 2, 2},
		{2, 3, 0},
		{4, 3, 1},
		{10, 3, 3},
		{10, -3, -3},
		{0, -3, 0},
		{-12, -3, 4},
		{-2147483648, -1, 2147483647},
	}

	for i, c := range cases {
		if fact := divide(c.dividend, c.divisor); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}

func Test(t *testing.T) {
	minusOne := int8(math.MinInt8)
	t.Logf("%b", uint8(minusOne))
}
