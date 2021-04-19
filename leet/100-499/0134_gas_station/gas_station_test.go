package _134_gas_station

import "testing"

func TestCanCompleterCircuit(t *testing.T) {
	var cases = []struct {
		gas, cost []int
		expected  int
	}{
		{
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			gas:      []int{1},
			cost:     []int{1},
			expected: 0,
		},
		{
			gas:      []int{1},
			cost:     []int{2},
			expected: -1,
		},
		{
			gas:      []int{1, 2},
			cost:     []int{2, 1},
			expected: 1,
		},
	}
	for i, c := range cases {
		if fact := canCompleteCircuit(c.gas, c.cost); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
