package _166_fraction_to_recurring_decimal

import (
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	cases := []struct {
		numerator, denominator int
		expected               string
	}{
		{1, 2, "0.5"},
		{0, 2, "0"},
		{0, -2, "0"},
		{2, 1, "2"},
		{1000, 333, "3.(003)"},
		{-1000, -333, "3.(003)"},
		{4, 9, "0.(4)"},
		{4, -9, "-0.(4)"},
		{-50, 8, "-6.25"},
		{7, -12, "-0.58(3)"},
	}

	for i, c := range cases {
		if fact := fractionToDecimal(c.numerator, c.denominator); fact != c.expected {
			t.Errorf("case#%d, want %s, got %s", i+1, c.expected, fact)
		}
	}
}
