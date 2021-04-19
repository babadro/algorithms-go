package _121_best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices   []int
		expected int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{prices: []int{7, 6, 4, 3, 1}, expected: 0},
		{prices: []int{1, 2, 3}, expected: 2},
		{prices: []int{}, expected: 0},
		{prices: []int{4}, expected: 0},
		{prices: []int{4, 1, 10, 2, 20, 1}, expected: 19},
		{prices: []int{2, 1, 2, 1, 0, 1, 2}, expected: 2},
		{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, expected: 4},
	}

	for i, c := range cases {
		if fact := maxProfit(c.prices); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
