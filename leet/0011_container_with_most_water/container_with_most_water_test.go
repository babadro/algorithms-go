package _011_container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 8, 6, 2, 5, 4, 8}, 40},
		{[]int{3, 2}, 2},
		{[]int{3, 2}, 2},
		{[]int{3, 5, 2, 5, 5, 2}, 15},
	}

	for i, c := range cases {
		if fact := maxArea(c.input); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
