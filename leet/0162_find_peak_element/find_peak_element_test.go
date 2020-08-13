package _162_find_peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 1, 3, 5, 4}, 0},
		{[]int{1, 4, 3, 6, 5}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 5},
		{[]int{3, 2, 1}, 0},
	}

	for i, c := range cases {
		if fact := findPeakElement(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
