package _033_search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{2, 1}, 1, 1},
		{[]int{3, 1, 2}, 2, 2},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1, 3}, 1, 0},
		{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{3, 1}, 3, 0},
		{[]int{3, 1, 2}, 3, 0},
	}

	for i, c := range cases {
		if fact := search3(c.nums, c.target); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
