package _136_single_number

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{7, 1, 6, 7, 1}, expected: 6},
	}

	for i, c := range cases {
		if fact := singleNumber(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
