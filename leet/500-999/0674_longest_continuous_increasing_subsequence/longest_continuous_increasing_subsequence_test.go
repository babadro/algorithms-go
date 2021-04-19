package _674_longest_continuous_increasing_subsequence

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 3}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{3, 2, 1}, 1},
		{[]int{2, 2, 2}, 1},
		{[]int{-1, 2, 1, 3, 5, 2, 3, 5, 6, 0}, 4},
		{[]int{3, 2, 1, 0, 1, 2, 3, 4}, 5},
	}

	for i, c := range cases {
		if fact := findLengthOfLCIS(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
