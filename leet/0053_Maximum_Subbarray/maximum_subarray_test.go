package _053_Maximum_Subbarray

import "testing"

func TestMaximumSubarray(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		//{[]int{2, 1, 3, 5, 4}, 15},
		//{[]int{3, 2, 4, 1, 5}, 15},
		//{[]int{-1, -2}, -1},
		//{[]int{-2, -1}, -1},
		//{[]int{-1}, -1},
		//{[]int{-2, 1}, 1},
		{[]int{1, -2}, 1},
		//{[]int{-1, 2,}, 2},
		//{[]int{2, -1,}, 2},
		//{[]int{-1, +1, -1, +1}, 1},
		//{[]int{1, -1, 1, -1}, 1},
		//{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for i, c := range cases {
		if fact := maxSubArray(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
