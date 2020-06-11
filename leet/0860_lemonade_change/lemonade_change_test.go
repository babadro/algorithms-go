package _860_lemonade_change

import "testing"

func TestLemonadeChange(t *testing.T) {
	cases := []struct {
		bills    []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{5}, true},
		{[]int{5, 10}, true},
		{[]int{10}, false},
		{[]int{20}, false},
		{[]int{10, 20}, false},
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 10, 5, 10, 5, 10}, true},
	}

	for i, c := range cases {
		if fact := lemonadeChange(c.bills); fact != c.expected {
			t.Errorf("case#%d: want %t, got %t", i+1, c.expected, fact)
		}
	}
}
