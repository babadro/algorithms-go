package _326_power_of_three

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		num      int
		expected bool
	}{
		{27, true},
		{0, false},
		{9, true},
		{45, false},
		{1, true},
		{-1, false},
		{3, true},
	}

	for i, c := range cases {
		if fact := isPowerOfThree(c.num); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
	}
}
